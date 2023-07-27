package main

import (
	_ "embed"
	"github.com/tmds-io/masterdata/service"

	"gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/cmd"
	"gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/service/telemetry/logger"
)

//go:generate container-gen

func main() {
	// We boot the service
	if err := service.Boot(); err != nil {
		logger.Fatal(err)
	}

	// Run service
	cmd.Run()
}
