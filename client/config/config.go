package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"gopkg.in/validator.v2"
)

type Config struct {
	UnderlayIP              string `json:"underlay_ip" validate:"nonzero"`
	SubnetPrefixLength      int    `json:"subnet_prefix_length" validate:"nonzero"`
	OverlayNetwork          string `json:"overlay_network" validate:"nonzero"`
	HealthCheckPort         uint16 `json:"health_check_port" validate:"nonzero"`
	VTEPName                string `json:"vtep_name" validate:"nonzero"`
	ConnectivityServerURL   string `json:"connectivity_server_url" validate:"nonzero"`
	ServerCACertFile        string `json:"ca_cert_file" validate:"nonzero"`
	ClientCertFile          string `json:"client_cert_file" validate:"nonzero"`
	ClientKeyFile           string `json:"client_key_file" validate:"nonzero"`
	VNI                     int    `json:"vni" validate:"nonzero"`
	PollInterval            int    `json:"poll_interval" validate:"nonzero"`
	DebugServerPort         int    `json:"debug_server_port" validate:"nonzero"`
	Datastore               string `json:"datastore" validate:"nonzero"`
	LeaseExpirationDuration int    `json:"lease_expiration_duration" validate:"nonzero"`
}

func LoadConfig(filePath string) (Config, error) {
	var cfg Config
	contents, err := ioutil.ReadFile(filePath)
	if err != nil {
		return cfg, fmt.Errorf("reading file %s: %s", filePath, err)
	}

	err = json.Unmarshal(contents, &cfg)
	if err != nil {
		return cfg, fmt.Errorf("unmarshaling contents: %s", err)
	}

	if err := validator.Validate(cfg); err != nil {
		return cfg, fmt.Errorf("invalid config: %s", err)
	}
	return cfg, nil
}
