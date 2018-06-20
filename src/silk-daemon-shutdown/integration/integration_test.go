package integration_test

import (
	"fmt"
	"os/exec"

	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
	"github.com/onsi/gomega/ghttp"
)

var (
	fakeRepServer         *ghttp.Server
	fakeSilkDaemonServer  *ghttp.Server
	tempPidFile           *os.File
	fakeSilkDaemonSession *gexec.Session
	tag                   string
)

const DEFAULT_TIMEOUT = "10s"

var _ = BeforeEach(func() {
	fakeRepServer = ghttp.NewUnstartedServer()
	fakeSilkDaemonServer = ghttp.NewUnstartedServer()
	fakeRepServer.AllowUnhandledRequests = true
	fakeSilkDaemonServer.AllowUnhandledRequests = true
	fakeRepServer.UnhandledRequestStatusCode = 500
	fakeSilkDaemonServer.UnhandledRequestStatusCode = 500

	fakeRepServer.AppendHandlers(ghttp.RespondWith(200, "", nil))
	fakeSilkDaemonServer.AppendHandlers(ghttp.RespondWith(200, "", nil))
})

var _ = JustBeforeEach(func() {
	fakeRepServer.Start()
	fakeSilkDaemonServer.Start()
})

var _ = AfterEach(func() {
	fakeRepServer.Close()
	fakeSilkDaemonServer.Close()

	if tempPidFile != nil {
		os.RemoveAll(tempPidFile.Name())
	}
})

var _ = Describe("Teardown", func() {
	AllIPTablesRules := func(tableName string) []string {
		iptablesSession, err := gexec.Start(exec.Command("iptables", "-w", "-S", "-t", tableName), GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())
		Eventually(iptablesSession).Should(gexec.Exit(0))
		return strings.Split(string(iptablesSession.Out.Contents()), "\n")
	}

	BeforeEach(func() {
		var err error
		tempPidFile, err = ioutil.TempFile(os.TempDir(), "pid")
		Expect(err).NotTo(HaveOccurred())
		sleepCommand := exec.Command("sleep", "60")

		fakeSilkDaemonSession, err = gexec.Start(sleepCommand, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())

		Expect(ioutil.WriteFile(tempPidFile.Name(), []byte(strconv.Itoa(sleepCommand.Process.Pid)+"\n"), 0777)).To(Succeed())
	})

	Context("when the servers eventually shutdown", func() {
		BeforeEach(func() {
			fakeRepServer.AppendHandlers(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
				go func() {
					fakeRepServer.Close()
				}()
			}))

			fakeSilkDaemonServer.AppendHandlers(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
				go func() {
					fakeSilkDaemonServer.Close()
				}()
			}))
		})

		It("kills the silk-daemon and pings the silk daemon until it stops responding", func() {
			session := runTeardown(fakeRepServer.URL(), fakeSilkDaemonServer.URL(), tempPidFile.Name())
			Eventually(session, DEFAULT_TIMEOUT).Should(gexec.Exit(0))

			Expect(fakeRepServer.ReceivedRequests()).To(HaveLen(2))
			Expect(session.Out).To(gbytes.Say("waiting for the rep to exit"))
			Expect(fakeRepServer.ReceivedRequests()).To(HaveLen(2))
			Expect(session.Out).To(gbytes.Say("sending TERM signal to silk-daemon"))
			Expect(session.Out).To(gbytes.Say("waiting for the silk daemon to exit"))
			Eventually(fakeSilkDaemonSession.ExitCode(), "5s").Should(Equal(143))
		})
	})

	Context("pinged servers return non 200 status codes", func() {
		It("pings the rep until the rep returns non 200 status code", func() {
			session := runTeardown(fakeRepServer.URL(), fakeSilkDaemonServer.URL(), tempPidFile.Name())
			Eventually(session, DEFAULT_TIMEOUT).Should(gexec.Exit(0))

			Expect(fakeRepServer.ReceivedRequests()).To(HaveLen(2))
			Expect(session.Out).To(gbytes.Say("waiting for the rep to exit"))
			Eventually(fakeSilkDaemonSession.ExitCode(), "5s").Should(Equal(143))
		})

		It("pings the silk daemon until it returns non 200 status code", func() {
			session := runTeardown(fakeRepServer.URL(), fakeSilkDaemonServer.URL(), tempPidFile.Name())
			Eventually(session, DEFAULT_TIMEOUT).Should(gexec.Exit(0))

			Expect(fakeRepServer.ReceivedRequests()).To(HaveLen(2))
			Expect(session.Out).To(gbytes.Say("waiting for the silk daemon to exit"))
			Eventually(fakeSilkDaemonSession.ExitCode(), "5s").Should(Equal(143))
		})
	})

	Context("when running in single ip mode", func() {
		BeforeEach(func() {
			iptablesSession, err := gexec.Start(exec.Command("iptables", "-N", "istio-ingress"), GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			Eventually(iptablesSession).Should(gexec.Exit(0))
			iptablesSession, err = gexec.Start(exec.Command("iptables", "-A", "OUTPUT", "-j", "istio-ingress"), GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			Eventually(iptablesSession).Should(gexec.Exit(0))
			iptablesSession, err = gexec.Start(exec.Command("iptables", "-A", "istio-ingress", "-o", "silk-vtep", "-j", "MARK", "--set-mark", "0"), GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			Eventually(iptablesSession).Should(gexec.Exit(0))
			iptablesSession, err = gexec.Start(exec.Command("iptables", "-A", "istio-ingress", "-o", "silk-vtep", "-j", "ACCEPT"), GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			Eventually(iptablesSession).Should(gexec.Exit(0))
		})

		It("deletes the iptables rule that marks overlay destined traffic", func() {
			session := runTeardown(fakeRepServer.URL(), fakeSilkDaemonServer.URL(), tempPidFile.Name())
			Eventually(session, DEFAULT_TIMEOUT).Should(gexec.Exit(0))

			rules := AllIPTablesRules("filter")
			Expect(rules).ToNot(ContainElement(ContainSubstring("-N istio-ingress")))
			Expect(rules).ToNot(ContainElement(ContainSubstring("-A OUTPUT -j istio-ingress")))
			Expect(rules).ToNot(ContainElement(ContainSubstring(fmt.Sprintf("-A istio-ingress -o silk-vtep -j MARK --set-xmark 0x%s/0xffffffff", tag))))
			Expect(rules).ToNot(ContainElement(ContainSubstring("-A istio-ingress -o silk-vtep -j ACCEPT")))
		})
	})

	Context("when connecting to the rep fails due to a bad url", func() {
		It("returns an error", func() {
			session := runTeardown("some/bad/url", fakeSilkDaemonServer.URL(), tempPidFile.Name())
			Eventually(session, DEFAULT_TIMEOUT).Should(gexec.Exit(1))

			Expect(session.Err).To(gbytes.Say("silk-daemon-shutdown: parse some/bad/url: invalid URI for request"))
		})
	})

	Context("when connecting to the silk-daemon fails due to a bad url", func() {
		It("returns an error", func() {
			session := runTeardown(fakeRepServer.URL(), "some/bad/url", tempPidFile.Name())
			Eventually(session, DEFAULT_TIMEOUT).Should(gexec.Exit(1))

			Expect(session.Err).To(gbytes.Say("silk-daemon-shutdown: parse some/bad/url: invalid URI for request"))
		})
	})

	Context("when pinging the rep takes a long time to reply", func() {
		BeforeEach(func() {
			fakeRepServer.AppendHandlers(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
				time.Sleep(10 * time.Second)
			}))
		})

		It("should timeout pinging rep and take less than 10 seconds to finish", func() {
			session := runTeardown(fakeRepServer.URL(), fakeSilkDaemonServer.URL(), tempPidFile.Name())
			Eventually(session, DEFAULT_TIMEOUT).Should(gexec.Exit(0))

			Expect(len(fakeRepServer.ReceivedRequests())).To(Equal(3))
			Expect(session.Out).To(gbytes.Say("pinging server timed out. trying again."))
			Expect(session.Out).To(gbytes.Say("waiting for the rep to exit"))
		})
	})

	Context("When silk daemon will not exit", func() {
		BeforeEach(func() {
			fakeSilkDaemonServer.UnhandledRequestStatusCode = 200
		})

		It("pings the silk daemon server 5 times and fails gracefully", func() {
			session := runTeardown(fakeRepServer.URL(), fakeSilkDaemonServer.URL(), tempPidFile.Name())
			Eventually(session, DEFAULT_TIMEOUT).Should(gexec.Exit(1))

			Expect(fakeSilkDaemonServer.ReceivedRequests()).To(HaveLen(5))
			Expect(session.Err).To(gbytes.Say("silk-daemon-shutdown: Silk Daemon Server did not exit after 5 ping attempts"))
		})
	})

	Context("When rep server will not exit", func() {
		BeforeEach(func() {
			fakeRepServer.UnhandledRequestStatusCode = 200
		})

		It("pings the rep server 40 times and fails gracefully", func() {
			session := runTeardown(fakeRepServer.URL(), fakeSilkDaemonServer.URL(), tempPidFile.Name())
			Eventually(session, DEFAULT_TIMEOUT).Should(gexec.Exit(0))

			Expect(fakeRepServer.ReceivedRequests()).To(HaveLen(40))
			Expect(session.Out).To(gbytes.Say("rep did not exit after 40 ping attempts"))

		}, 5)
	})

	Context("when silk daemon pid file does not exist", func() {
		It("returns an error", func() {
			session := runTeardown(fakeRepServer.URL(), fakeSilkDaemonServer.URL(), "/some-invalid/file-path")
			Eventually(session, DEFAULT_TIMEOUT).Should(gexec.Exit(1))

			Expect(session.Err).To(gbytes.Say("silk-daemon-shutdown: open /some-invalid/file-path: no such file or directory"))
		})
	})

	Context("when the silk daemon is not running", func() {
		BeforeEach(func() {
			fakeSilkDaemonSession.Kill()
			Eventually(fakeSilkDaemonSession).Should(gexec.Exit())
		})

		It("returns with exit code 0", func() {
			session := runTeardown(fakeRepServer.URL(), fakeSilkDaemonServer.URL(), tempPidFile.Name())
			Eventually(session, DEFAULT_TIMEOUT).Should(gexec.Exit(0))
		})
	})

	Context("when silk daemon pid file does not contain a number", func() {
		BeforeEach(func() {
			Expect(ioutil.WriteFile(tempPidFile.Name(), []byte("not-a-number"), 0777)).To(Succeed())
		})

		It("returns an error", func() {
			session := runTeardown(fakeRepServer.URL(), fakeSilkDaemonServer.URL(), tempPidFile.Name())
			Eventually(session, DEFAULT_TIMEOUT).Should(gexec.Exit(1))

			Expect(session.Err).To(gbytes.Say("silk-daemon-shutdown: strconv.Atoi: parsing \"not-a-number\": invalid syntax"))
		})
	})
})

func runTeardown(url, silkDaemonUrl, silkDaemonPidFile string) *gexec.Session {
	startCmd := exec.Command(paths.TeardownBin,
		"--repUrl", url,
		"--silkDaemonUrl", silkDaemonUrl,
		"--repTimeout", "0",
		"--silkDaemonTimeout", "0",
		"--silkDaemonPidPath", silkDaemonPidFile,
		"--iptablesLockFile", "/tmp/someLockWhoReallyCares.lock")
	session, err := gexec.Start(startCmd, GinkgoWriter, GinkgoWriter)
	Expect(err).NotTo(HaveOccurred())
	return session
}
