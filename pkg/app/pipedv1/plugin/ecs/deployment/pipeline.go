package deployment

import (
	sdk "github.com/pipe-cd/piped-plugin-sdk-go"
)

const (
	// StageECSSync represents the ECS sync stage.
	StageECSSync = "ECS_SYNC"

	// StageECSPrimaryRollout represents the ECS primary rollout stage.
	StageECSPrimaryRollout = "ECS_PRIMARY_ROLLOUT"

	// StageECSCanaryRollout represents the ECS canary rollout stage.
	StageECSCanaryRollout = "ECS_CANARY_ROLLOUT"

	// StageECSCanaryClean represents the ECS canary clean stage.
	StageECSCanaryClean = "ECS_CANARY_CLEAN"

	// StageECSTrafficRouting represents the ECS traffic routing stage.
	StageECSTrafficRouting = "ECS_TRAFFIC_ROUTING"

	// StageECSRollback represents the ECS rollback stage.
	StageECSRollback = "ECS_ROLLBACK"
)

var allStages = []string{
	StageECSSync,
	StageECSPrimaryRollout,
	StageECSCanaryRollout,
	StageECSCanaryClean,
	StageECSTrafficRouting,
	StageECSRollback,
}

const (
	StageECSSyncDescription           = "Sync ECS service with given task definition"
	StageECSPrimaryRolloutDescription = "Roll out new task set as primary"
	StageECSCanaryRolloutDescription  = "Roll out new task set as canary"
	StageECSCanaryCleanDescription    = "Clean up canary task set"
	StageECSTrafficRoutingDescription = "Route traffic between primary and canary task sets"
	StageECSRollbackDescription       = "Rollback to previous task set"
)

func buildQuickSyncPipeline(autoRollback bool) []sdk.QuickSyncStage {
	out := []sdk.QuickSyncStage{
		{
			Name:        StageECSSync,
			Description: StageECSSyncDescription,
			Rollback:    false,
		},
	}

	if autoRollback {
		out = append(out, sdk.QuickSyncStage{
			Name:        StageECSRollback,
			Description: StageECSRollbackDescription,
			Rollback:    true,
		})
	}

	return out
}

func buildPipelineStages(input *sdk.BuildPipelineSyncStagesInput) []sdk.PipelineStage {
	stages := input.Request.Stages

	out := make([]sdk.PipelineStage, 0, len(stages)+1)

	for _, s := range stages {
		out = append(out, sdk.PipelineStage{
			Name:     s.Name,
			Index:    s.Index,
			Rollback: false,
		})
	}

	if input.Request.Rollback {
		out = append(out, sdk.PipelineStage{
			Name:     StageECSRollback,
			Index:    len(stages) + 1,
			Rollback: true,
		})
	}

	return out
}
