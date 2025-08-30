package config

import (
	"fmt"

	"github.com/jnnkrdb/orasul/pkg/envconfig"
	"github.com/jnnkrdb/orasul/pkg/logging"
)

type Config struct {
	OCI struct {
		Registry string
		Username string
		Password string
	}
}

var Cfg Config

func LoadConfig() {
	if err := envconfig.Process("orasul", &Cfg); err != nil {
		fmt.Printf("Error loading config: %v\n", err)
	}

	logging.Default.Debug("starting with config: %v", Cfg)
}
