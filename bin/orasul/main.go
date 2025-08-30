package main

import (
	"time"

	"github.com/jnnkrdb/orasul/bin/orasul/config"
	"github.com/jnnkrdb/orasul/pkg/logging"
)

func main() {

	config.LoadConfig()

	var round int = 1
	for {
		time.Sleep(5 * time.Minute)
		logging.Default.Debug("Current round %d", round)
		round++
	}
}
