require 'rspec'
require 'bosh/template/test'
require 'yaml'
require 'json'


module Bosh::Template::Test
  describe 'vxlan-policy-agent job' do
    let(:merged_manifest_properties) do
      {
        "log_level" => 'high',
        "log_prefix" => "cfnetworking",
        "iptables_logging" => true,
        "iptables_accepted_udp_logs_per_sec" => 5,
        "poll_interval" => 5,
        "policy_server" => {
          "hostname" => "example.internal",
          "internal_listen_port" => 4002,
        },
        "metron_port" => 5000,
        "debug_server_port" => 1010,
        "disable_container_network_policy" => false,
     }
    end

      links = [
        Link.new(
          name: 'cf_network',
          instances: [LinkInstance.new()],
          properties: {
            'network' => '10.255.0.0/16',
          }
        )
      ]

    let(:release_path) { File.join(File.dirname(__FILE__), '../..') }
    let(:release) {ReleaseDir.new(release_path)}
    let(:job) {release.job('vxlan-policy-agent')}

    describe 'vxlan-policy-agent spec' do
      let (:template) { job.template('config/vxlan-policy-agent.json') }

      it 'does something' do
        config = JSON.parse(template.render(merged_manifest_properties, consumes: links))
        expect(config).to eq({
          "log_level" => 'high',
          "log_prefix" => "cfnetworking",
          "iptables_c2c_logging" => true,
          "iptables_accepted_udp_logs_per_sec" => 5,
          "poll_interval" => 5,

          "policy_server_url" => "https://example.internal:4002",
          "metron_address" => "127.0.0.1:5000",
          "debug_server_port" => 1010,
          "disable_container_network_policy" => false,
          'overlay_network' => '10.255.0.0/16',

          "ca_cert_file" => "/var/vcap/jobs/vxlan-policy-agent/config/certs/ca.crt",
          "client_cert_file" => "/var/vcap/jobs/vxlan-policy-agent/config/certs/client.crt",
          "client_key_file" => "/var/vcap/jobs/vxlan-policy-agent/config/certs/client.key",

          "cni_datastore_path" => "/var/vcap/data/container-metadata/store.json",
          "iptables_lock_file" => "/var/vcap/data/garden-cni/iptables.lock",
          "debug_server_host" => "127.0.0.1",
          "client_timeout_seconds" => 5,
          "vni" => 1,
        })

      end
    end
  end
end

