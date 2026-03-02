package provider

import (
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/pipe-cd/pipecd/pkg/app/pipedv1/plugin/ecs/config"
)

var ErrNoTargetGroup = errors.New("no target group")

// loadTargetGroups returns the primary and canary target groups from the given config.
func loadTargetGroups(targetGroups config.ECSTargetGroups) (*types.LoadBalancer, *types.LoadBalancer, error) {
	if targetGroups.Primary == nil {
		return nil, nil, ErrNoTargetGroup
	}

	primary := &types.LoadBalancer{
		TargetGroupArn: aws.String(targetGroups.Primary.TargetGroupARN),
		ContainerName:  aws.String(targetGroups.Primary.ContainerName),
		ContainerPort:  aws.Int32(targetGroups.Primary.ContainerPort),
	}

	var canary *types.LoadBalancer
	if targetGroups.Canary != nil {
		canary = &types.LoadBalancer{
			TargetGroupArn: aws.String(targetGroups.Canary.TargetGroupARN),
			ContainerName:  aws.String(targetGroups.Canary.ContainerName),
			ContainerPort:  aws.Int32(targetGroups.Canary.ContainerPort),
		}
	}

	return primary, canary, nil
}
