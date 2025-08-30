package config

import (
	"fmt"

	"github.com/jnnkrdb/orasul/pkg/envconfig"
	"github.com/jnnkrdb/orasul/pkg/logging"
)

type Config struct {
	OCI struct {
		Registry string `env:"ORASUL_OCI_REGISTRY" default:""`
		Username string `env:"ORASUL_OCI_USERNAME" default:""`
		Password string `env:"ORASUL_OCI_PASSWORD" default:""`
	}
}

var Cfg Config

func LoadConfig() {
	if err := envconfig.Process("", &Cfg); err != nil {
		fmt.Printf("Error loading config: %v\n", err)
	}

	logging.Default.Debug("starting with config: %v", Cfg)
}
