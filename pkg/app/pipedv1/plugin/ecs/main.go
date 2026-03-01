package main

import (
	"log"

	"github.com/pipe-cd/pipecd/pkg/app/pipedv1/plugin/ecs/deployment"
	sdk "github.com/pipe-cd/piped-plugin-sdk-go"
)

func main() {
	plugin, err := sdk.NewPlugin(
		"0.0.1",
		sdk.WithDeploymentPlugin(&deployment.ECSPlugin{}),
	)
	if err != nil {
		log.Fatalf("failed to create plugin: %v", err)
	}
	if err := plugin.Run(); err != nil {
		log.Fatalf("plugin execution failed: %v", err)
	}
}
