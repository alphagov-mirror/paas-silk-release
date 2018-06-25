package rules_test

import (
	"errors"
	"fmt"
	"lib/fakes"
	"lib/rules"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("LockedIptables", func() {
	var (
		lockedIPT *rules.LockedIPTables
		ipt       *fakes.IPTables
		restorer  *fakes.Restorer
		lock      *fakes.Locker
		rulespec  []string
		rule      rules.IPTablesRule
	)
	BeforeEach(func() {
		ipt = &fakes.IPTables{}
		lock = &fakes.Locker{}
		restorer = &fakes.Restorer{}
		lockedIPT = &rules.LockedIPTables{
			IPTables: ipt,
			Locker:   lock,
			Restorer: restorer,
		}
		rulespec = []string{"some", "args"}
		rule = rules.IPTablesRule{"some", "args"}
	})
	Describe("BulkInsert", func() {
		var ruleSet []rules.IPTablesRule
		BeforeEach(func() {
			ruleSet = []rules.IPTablesRule{
				rules.NewMarkSetRule("1.2.3.4", "A", "a-guid"),
				rules.NewMarkSetRule("2.2.2.2", "B", "b-guid"),
			}
		})

		It("constructs the input and passes it to the restorer", func() {
			err := lockedIPT.BulkInsert("some-table", "some-chain", 1, ruleSet...)
			Expect(err).NotTo(HaveOccurred())

			Expect(lock.LockCallCount()).To(Equal(1))
			Expect(lock.UnlockCallCount()).To(Equal(1))
			Expect(restorer.RestoreCallCount()).To(Equal(1))
			restoreInput := restorer.RestoreArgsForCall(0)
			Expect(restoreInput).To(ContainSubstring("*some-table\n"))
			Expect(restoreInput).To(ContainSubstring("-I some-chain 1 --source 1.2.3.4 --jump MARK --set-xmark 0xA -m comment --comment src:a-guid\n"))
			Expect(restoreInput).To(ContainSubstring("-I some-chain 1 --source 2.2.2.2 --jump MARK --set-xmark 0xB -m comment --comment src:b-guid\n"))
			Expect(restoreInput).To(ContainSubstring("COMMIT\n"))
		})
		Context("when the lock fails", func() {
			BeforeEach(func() {
				lock.LockReturns(errors.New("banana"))
			})
			It("should return an error", func() {
				err := lockedIPT.BulkInsert("some-table", "some-chain", 1, ruleSet...)
				Expect(err).To(MatchError("lock: banana"))
			})
		})

		Context("when the restorer fails", func() {
			BeforeEach(func() {
				restorer.RestoreReturns(fmt.Errorf("banana"))
			})
			It("should return an error", func() {
				err := lockedIPT.BulkInsert("some-table", "some-chain", 1, ruleSet...)
				Expect(err).To(MatchError("iptables call: banana and unlock: <nil>"))
			})
		})

		Context("when the unlock fails", func() {
			BeforeEach(func() {
				lock.UnlockReturns(errors.New("banana"))
			})
			It("should return an error", func() {
				err := lockedIPT.BulkInsert("some-table", "some-chain", 1, ruleSet...)
				Expect(err).To(MatchError("banana"))
			})
		})
		Context("when the restorer fails and then the unlock fails", func() {
			BeforeEach(func() {
				lock.UnlockReturns(errors.New("banana"))
				restorer.RestoreReturns(fmt.Errorf("patato"))
			})
			It("should return an error", func() {
				err := lockedIPT.BulkInsert("some-table", "some-chain", 1, ruleSet...)
				Expect(err).To(MatchError("iptables call: patato and unlock: banana"))
			})
		})
	})

	Describe("BulkAppend", func() {
		var ruleSet []rules.IPTablesRule
		BeforeEach(func() {
			ruleSet = []rules.IPTablesRule{
				rules.NewMarkSetRule("1.2.3.4", "A", "a-guid"),
				rules.NewMarkSetRule("2.2.2.2", "B", "b-guid"),
			}
		})

		It("constructs the input and passes it to the restorer", func() {
			err := lockedIPT.BulkAppend("some-table", "some-chain", ruleSet...)
			Expect(err).NotTo(HaveOccurred())

			Expect(lock.LockCallCount()).To(Equal(1))
			Expect(lock.UnlockCallCount()).To(Equal(1))
			Expect(restorer.RestoreCallCount()).To(Equal(1))
			restoreInput := restorer.RestoreArgsForCall(0)
			Expect(restoreInput).To(ContainSubstring("*some-table\n"))
			Expect(restoreInput).To(ContainSubstring("-A some-chain --source 1.2.3.4 --jump MARK --set-xmark 0xA -m comment --comment src:a-guid\n"))
			Expect(restoreInput).To(ContainSubstring("-A some-chain --source 2.2.2.2 --jump MARK --set-xmark 0xB -m comment --comment src:b-guid\n"))
			Expect(restoreInput).To(ContainSubstring("COMMIT\n"))
		})
		Context("when the lock fails", func() {
			BeforeEach(func() {
				lock.LockReturns(errors.New("banana"))
			})
			It("should return an error", func() {
				err := lockedIPT.BulkAppend("some-table", "some-chain", ruleSet...)
				Expect(err).To(MatchError("lock: banana"))
			})
		})

		Context("when the restorer fails", func() {
			BeforeEach(func() {
				restorer.RestoreReturns(fmt.Errorf("banana"))
			})
			It("should return an error", func() {
				err := lockedIPT.BulkAppend("some-table", "some-chain", ruleSet...)
				Expect(err).To(MatchError("iptables call: banana and unlock: <nil>"))
			})
		})

		Context("when the unlock fails", func() {
			BeforeEach(func() {
				lock.UnlockReturns(errors.New("banana"))
			})
			It("should return an error", func() {
				err := lockedIPT.BulkAppend("some-table", "some-chain", ruleSet...)
				Expect(err).To(MatchError("banana"))
			})
		})
		Context("when the restorer fails and then the unlock fails", func() {
			BeforeEach(func() {
				lock.UnlockReturns(errors.New("banana"))
				restorer.RestoreReturns(fmt.Errorf("patato"))
			})
			It("should return an error", func() {
				err := lockedIPT.BulkAppend("some-table", "some-chain", ruleSet...)
				Expect(err).To(MatchError("iptables call: patato and unlock: banana"))
			})
		})
	})

	Describe("Exists", func() {
		BeforeEach(func() {
			ipt.ExistsReturns(true, nil)
		})
		It("passes the correct parameters to the iptables library", func() {
			exists, err := lockedIPT.Exists("some-table", "some-chain", rule)
			Expect(err).NotTo(HaveOccurred())
			Expect(exists).To(Equal(true))

			Expect(lock.LockCallCount()).To(Equal(1))
			Expect(lock.UnlockCallCount()).To(Equal(1))
			Expect(ipt.ExistsCallCount()).To(Equal(1))
			table, chain, spec := ipt.ExistsArgsForCall(0)
			Expect(table).To(Equal("some-table"))
			Expect(chain).To(Equal("some-chain"))
			Expect(spec).To(Equal(rulespec))
		})

		Context("when locking fails", func() {
			BeforeEach(func() {
				lock.LockReturns(errors.New("banana"))
			})
			It("returns an error", func() {
				_, err := lockedIPT.Exists("some-table", "some-chain", rule)
				Expect(err).To(MatchError("lock: banana"))
			})
		})

		Context("when iptables call fails and unlock succeeds", func() {
			BeforeEach(func() {
				ipt.ExistsReturns(false, errors.New("banana"))
			})
			It("returns an error", func() {
				_, err := lockedIPT.Exists("some-table", "some-chain", rule)
				Expect(err).To(MatchError("iptables call: banana and unlock: <nil>"))
			})
		})

		Context("when iptables call fails and unlock fails", func() {
			BeforeEach(func() {
				lock.UnlockReturns(errors.New("banana"))
				ipt.ExistsReturns(false, errors.New("patato"))
			})
			It("returns an error", func() {
				_, err := lockedIPT.Exists("some-table", "some-chain", rule)
				Expect(err).To(MatchError("iptables call: patato and unlock: banana"))
			})
		})
	})

	Describe("Delete", func() {
		It("locks and passes the correct parameters to the iptables library", func() {
			err := lockedIPT.Delete("some-table", "some-chain", rule)
			Expect(err).NotTo(HaveOccurred())

			Expect(lock.LockCallCount()).To(Equal(1))
			Expect(lock.UnlockCallCount()).To(Equal(1))
			Expect(ipt.DeleteCallCount()).To(Equal(1))
			table, chain, spec := ipt.DeleteArgsForCall(0)
			Expect(table).To(Equal("some-table"))
			Expect(chain).To(Equal("some-chain"))
			Expect(spec).To(Equal(rulespec))
		})

		Context("when locking fails", func() {
			BeforeEach(func() {
				lock.LockReturns(errors.New("banana"))
			})
			It("returns an error", func() {
				err := lockedIPT.Delete("some-table", "some-chain", rule)
				Expect(err).To(MatchError("lock: banana"))
			})
		})

		Context("when iptables call fails and unlock succeeds", func() {
			BeforeEach(func() {
				ipt.DeleteReturns(errors.New("banana"))
			})
			It("returns an error", func() {
				err := lockedIPT.Delete("some-table", "some-chain", rule)
				Expect(err).To(MatchError("iptables call: banana and unlock: <nil>"))
			})
		})

		Context("when iptables call fails and unlock fails", func() {
			BeforeEach(func() {
				lock.UnlockReturns(errors.New("banana"))
				ipt.DeleteReturns(errors.New("patato"))
			})
			It("returns an error", func() {
				err := lockedIPT.Delete("some-table", "some-chain", rule)
				Expect(err).To(MatchError("iptables call: patato and unlock: banana"))
			})
		})
	})

	Describe("List", func() {
		BeforeEach(func() {
			ipt.ListReturns([]string{"some", "list"}, nil)
		})
		It("locks and passes the correct parameters to the iptables library", func() {
			list, err := lockedIPT.List("some-table", "some-chain")
			Expect(err).NotTo(HaveOccurred())
			Expect(list).To(Equal([]string{"some", "list"}))

			Expect(lock.LockCallCount()).To(Equal(1))
			Expect(lock.UnlockCallCount()).To(Equal(1))
			Expect(ipt.ListCallCount()).To(Equal(1))
			table, chain := ipt.ListArgsForCall(0)
			Expect(table).To(Equal("some-table"))
			Expect(chain).To(Equal("some-chain"))
		})

		Context("when locking fails", func() {
			BeforeEach(func() {
				lock.LockReturns(errors.New("banana"))
			})
			It("returns an error", func() {
				_, err := lockedIPT.List("some-table", "some-chain")
				Expect(err).To(MatchError("lock: banana"))
			})
		})

		Context("when iptables call fails and unlock succeeds", func() {
			BeforeEach(func() {
				ipt.ListReturns(nil, errors.New("banana"))
			})
			It("returns an error", func() {
				_, err := lockedIPT.List("some-table", "some-chain")
				Expect(err).To(MatchError("iptables call: banana and unlock: <nil>"))
			})
		})

		Context("when iptables call fails and unlock fails", func() {
			BeforeEach(func() {
				lock.UnlockReturns(errors.New("banana"))
				ipt.ListReturns(nil, errors.New("patato"))
			})
			It("returns an error", func() {
				_, err := lockedIPT.List("some-table", "some-chain")
				Expect(err).To(MatchError("iptables call: patato and unlock: banana"))
			})
		})
	})

	Describe("NewChain", func() {
		It("locks and passes the correct parameters to the iptables library", func() {
			err := lockedIPT.NewChain("some-table", "some-chain")
			Expect(err).NotTo(HaveOccurred())

			Expect(lock.LockCallCount()).To(Equal(1))
			Expect(lock.UnlockCallCount()).To(Equal(1))
			Expect(ipt.NewChainCallCount()).To(Equal(1))
			table, chain := ipt.NewChainArgsForCall(0)
			Expect(table).To(Equal("some-table"))
			Expect(chain).To(Equal("some-chain"))
		})

		Context("when locking fails", func() {
			BeforeEach(func() {
				lock.LockReturns(errors.New("banana"))
			})
			It("returns an error", func() {
				err := lockedIPT.NewChain("some-table", "some-chain")
				Expect(err).To(MatchError("lock: banana"))
			})
		})

		Context("when iptables call fails and unlock succeeds", func() {
			BeforeEach(func() {
				ipt.NewChainReturns(errors.New("banana"))
			})
			It("returns an error", func() {
				err := lockedIPT.NewChain("some-table", "some-chain")
				Expect(err).To(MatchError("iptables call: banana and unlock: <nil>"))
			})
		})

		Context("when iptables call fails and unlock fails", func() {
			BeforeEach(func() {
				lock.UnlockReturns(errors.New("banana"))
				ipt.NewChainReturns(errors.New("patato"))
			})
			It("returns an error", func() {
				err := lockedIPT.NewChain("some-table", "some-chain")
				Expect(err).To(MatchError("iptables call: patato and unlock: banana"))
			})
		})
	})

	Describe("DeleteChain", func() {
		It("locks and passes the correct parameters to the iptables library", func() {
			err := lockedIPT.DeleteChain("some-table", "some-chain")
			Expect(err).NotTo(HaveOccurred())

			Expect(lock.LockCallCount()).To(Equal(1))
			Expect(lock.UnlockCallCount()).To(Equal(1))
			Expect(ipt.DeleteChainCallCount()).To(Equal(1))
			table, chain := ipt.DeleteChainArgsForCall(0)
			Expect(table).To(Equal("some-table"))
			Expect(chain).To(Equal("some-chain"))
		})

		Context("when locking fails", func() {
			BeforeEach(func() {
				lock.LockReturns(errors.New("banana"))
			})
			It("returns an error", func() {
				err := lockedIPT.DeleteChain("some-table", "some-chain")
				Expect(err).To(MatchError("lock: banana"))
			})
		})

		Context("when iptables call fails and unlock succeeds", func() {
			BeforeEach(func() {
				ipt.DeleteChainReturns(errors.New("banana"))
			})
			It("returns an error", func() {
				err := lockedIPT.DeleteChain("some-table", "some-chain")
				Expect(err).To(MatchError("iptables call: banana and unlock: <nil>"))
			})
		})

		Context("when iptables call fails and unlock fails", func() {
			BeforeEach(func() {
				lock.UnlockReturns(errors.New("banana"))
				ipt.DeleteChainReturns(errors.New("patato"))
			})
			It("returns an error", func() {
				err := lockedIPT.DeleteChain("some-table", "some-chain")
				Expect(err).To(MatchError("iptables call: patato and unlock: banana"))
			})
		})
	})

	Describe("ClearChain", func() {
		It("locks and passes the correct parameters to the iptables library", func() {
			err := lockedIPT.ClearChain("some-table", "some-chain")
			Expect(err).NotTo(HaveOccurred())

			Expect(lock.LockCallCount()).To(Equal(1))
			Expect(lock.UnlockCallCount()).To(Equal(1))
			Expect(ipt.ClearChainCallCount()).To(Equal(1))
			table, chain := ipt.ClearChainArgsForCall(0)
			Expect(table).To(Equal("some-table"))
			Expect(chain).To(Equal("some-chain"))
		})

		Context("when locking fails", func() {
			BeforeEach(func() {
				lock.LockReturns(errors.New("banana"))
			})
			It("returns an error", func() {
				err := lockedIPT.ClearChain("some-table", "some-chain")
				Expect(err).To(MatchError("lock: banana"))
			})
		})

		Context("when iptables call fails and unlock succeeds", func() {
			BeforeEach(func() {
				ipt.ClearChainReturns(errors.New("banana"))
			})
			It("returns an error", func() {
				err := lockedIPT.ClearChain("some-table", "some-chain")
				Expect(err).To(MatchError("iptables call: banana and unlock: <nil>"))
			})
		})

		Context("when iptables call fails and unlock fails", func() {
			BeforeEach(func() {
				lock.UnlockReturns(errors.New("banana"))
				ipt.ClearChainReturns(errors.New("patato"))
			})
			It("returns an error", func() {
				err := lockedIPT.ClearChain("some-table", "some-chain")
				Expect(err).To(MatchError("iptables call: patato and unlock: banana"))
			})
		})
	})

	Describe("ListAll", func() {
		BeforeEach(func() {
			ipt.ListChainsReturns([]string{
				"a chain",
				"another chain",
				"a third chain",
			}, nil)

			ipt.ListReturnsOnCall(0, []string{
				"a rule",
				"another rule",
				"a third rule",
			}, nil)

			ipt.ListReturnsOnCall(1, []string{
				"a rule",
			}, nil)

			ipt.ListReturnsOnCall(2, []string{
				"a rule",
				"another rule",
			}, nil)
		})
		It("should return all the rows", func() {
			rows, err := lockedIPT.ListAll("some-table")
			Expect(err).NotTo(HaveOccurred())

			Expect(lock.LockCallCount()).To(Equal(1))
			Expect(lock.UnlockCallCount()).To(Equal(1))
			Expect(ipt.ListChainsCallCount()).To(Equal(1))
			table := ipt.ListChainsArgsForCall(0)
			Expect(table).To(Equal("some-table"))

			Expect(ipt.ListCallCount()).To(Equal(3))
			table, chain := ipt.ListArgsForCall(0)
			Expect(table).To(Equal("some-table"))
			Expect(chain).To(Equal("a chain"))
			table, chain = ipt.ListArgsForCall(1)
			Expect(table).To(Equal("some-table"))
			Expect(chain).To(Equal("another chain"))
			table, chain = ipt.ListArgsForCall(2)
			Expect(table).To(Equal("some-table"))
			Expect(chain).To(Equal("a third chain"))

			Expect(rows).To(Equal([]string{
				"a rule",
				"another rule",
				"a third rule",
				"a rule",
				"a rule",
				"another rule",
			}))
		})

		Context("when locking fails", func() {
			BeforeEach(func() {
				lock.LockReturns(errors.New("banana"))
			})
			It("returns an error", func() {
				_, err := lockedIPT.ListAll("some-table")
				Expect(err).To(MatchError("lock: banana"))
			})
		})

		Context("when list chains call fails and unlock succeeds", func() {
			BeforeEach(func() {
				ipt.ListChainsReturns(nil, errors.New("fig"))
			})
			It("returns an error", func() {
				_, err := lockedIPT.ListAll("some-table")
				Expect(err).To(MatchError("iptables call: fig and unlock: <nil>"))
			})
		})

		Context("when list chains call fails and unlock fails", func() {
			BeforeEach(func() {
				lock.UnlockReturns(errors.New("apple"))
				ipt.ListChainsReturns(nil, errors.New("patato"))
			})
			It("returns an error", func() {
				_, err := lockedIPT.ListAll("some-table")
				Expect(err).To(MatchError("iptables call: patato and unlock: apple"))
			})
		})

		Context("when list call fails and unlock succeeds", func() {
			BeforeEach(func() {
				ipt.ListReturnsOnCall(0, nil, errors.New("kumquat"))
			})
			It("returns an error", func() {
				_, err := lockedIPT.ListAll("some-table")
				Expect(err).To(MatchError("iptables call: kumquat and unlock: <nil>"))
			})
		})

		Context("when list call fails and unlock fails", func() {
			BeforeEach(func() {
				lock.UnlockReturns(errors.New("cherry"))
				ipt.ListReturnsOnCall(0, nil, errors.New("avacado"))
			})
			It("returns an error", func() {
				_, err := lockedIPT.ListAll("some-table")
				Expect(err).To(MatchError("iptables call: avacado and unlock: cherry"))
			})
		})
	})
})
