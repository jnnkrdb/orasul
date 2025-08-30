package config

import (
	"github.com/jnnkrdb/orasul/pkg/envconfig"
	"github.com/jnnkrdb/orasul/pkg/logging"
)

type Config struct {
	Oci struct {
		Registry string
		Username string
		Password string
	}
	Local struct {
		RegistryPath string `default:"/opt/orasul/data"`
	}
}

var Cfg Config

func LoadConfig() {
	if err := envconfig.Process("orasul", &Cfg); err != nil {
		logging.Default.Error("Error loading config", "err", err)
	}

	logging.Default.Debug("starting with config", "config", Cfg)
}
