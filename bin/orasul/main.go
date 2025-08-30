package main

import (
	"fmt"
	"time"

	"github.com/jnnkrdb/orasul/bin/orasul/config"
	"github.com/jnnkrdb/orasul/bin/orasul/oci"
	"github.com/jnnkrdb/orasul/pkg/logging"
)

const _MEDIATYPE = "application/vnd.oci.artifact.manifest.v1+json"

func main() {

	config.LoadConfig()

	// upload file
	reg := fmt.Sprintf("%s/%s:%s", config.Cfg.Oci.Registry, "backups/test", "latest")
	oci.UploadToRegistry(reg, _MEDIATYPE, "/opt/orasul/data/test.json")

	var round int = 1
	for {
		time.Sleep(5 * time.Minute)
		logging.Default.Debug("Current round %d", round)
		round++
	}
}
