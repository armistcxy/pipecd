package deployment

import (
	"context"

	sdk "github.com/pipe-cd/piped-plugin-sdk-go"

	ecsconfig "github.com/pipe-cd/pipecd/pkg/app/pipedv1/plugin/ecs/config"
)

var _ sdk.DeploymentPlugin[ecsconfig.ECSPluginConfig, ecsconfig.ECSDeployTargetConfig, ecsconfig.ECSApplicationSpec] = (*ECSPlugin)(nil)

type ECSPlugin struct {
}

// FetchDefinedStages returns the list of stages that the plugin can execute.
func (p *ECSPlugin) FetchDefinedStages() []string {
	return allStages
}

// BuildQuickSyncStages builds the stages that will be executed during the quick sync process.
func (p *ECSPlugin) BuildQuickSyncStages(
	ctx context.Context,
	cfg *ecsconfig.ECSPluginConfig,
	input *sdk.BuildQuickSyncStagesInput,
) (*sdk.BuildQuickSyncStagesResponse, error) {
	panic("unimplemented")
}

// BuildPipelineSyncStages builds the stages that will be executed by the plugin.
func (p *ECSPlugin) BuildPipelineSyncStages(
	_ context.Context,
	_ *ecsconfig.ECSPluginConfig,
	input *sdk.BuildPipelineSyncStagesInput,
) (*sdk.BuildPipelineSyncStagesResponse, error) {
	panic("unimplemented")
}

// ExecuteStage executes the given stage.
func (p *ECSPlugin) ExecuteStage(
	ctx context.Context,
	cfg *ecsconfig.ECSPluginConfig,
	deployTargets []*sdk.DeployTarget[ecsconfig.ECSDeployTargetConfig],
	input *sdk.ExecuteStageInput[ecsconfig.ECSApplicationSpec],
) (*sdk.ExecuteStageResponse, error) {
	panic("unimplemented")
}

// DetermineVersions determines the versions of the resources that will be deployed.
func (p *ECSPlugin) DetermineVersions(
	ctx context.Context,
	cfg *ecsconfig.ECSPluginConfig,
	input *sdk.DetermineVersionsInput[ecsconfig.ECSApplicationSpec],
) (*sdk.DetermineVersionsResponse, error) {
	panic("unimplemented")
}

// DetermineStrategy determines the strategy to deploy the resources.
func (p *ECSPlugin) DetermineStrategy(
	ctx context.Context,
	cfg *ecsconfig.ECSPluginConfig,
	input *sdk.DetermineStrategyInput[ecsconfig.ECSApplicationSpec],
) (*sdk.DetermineStrategyResponse, error) {
	panic("unimplemented")
}
