package provider

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/go-playground/assert/v2"
	"github.com/pipe-cd/pipecd/pkg/app/pipedv1/plugin/ecs/config"
)

func TestLoadTargetGroups(t *testing.T) {
	t.Parallel()

	testcases := []struct {
		name        string
		cfg         config.ECSTargetGroups
		expected    []*types.LoadBalancer
		expectedErr bool
	}{
		{
			name:        "no target group",
			cfg:         config.ECSTargetGroups{},
			expected:    []*types.LoadBalancer{nil, nil},
			expectedErr: true,
		},
		{
			name: "primary target group only",
			cfg: config.ECSTargetGroups{
				Primary: &config.ECSTargetGroup{
					TargetGroupARN: "primary-target-group-arn",
					ContainerName:  "primary-container-name",
					ContainerPort:  80,
				},
			},
			expected: []*types.LoadBalancer{
				{
					TargetGroupArn: aws.String("primary-target-group-arn"),
					ContainerName:  aws.String("primary-container-name"),
					ContainerPort:  aws.Int32(80),
				},
				nil,
			},
			expectedErr: false,
		},
		{
			name: "primary and canary target group",
			cfg: config.ECSTargetGroups{
				Primary: &config.ECSTargetGroup{
					TargetGroupARN: "primary-target-group-arn",
					ContainerName:  "primary-container-name",
					ContainerPort:  80,
				},
				Canary: &config.ECSTargetGroup{
					TargetGroupARN: "canary-target-group-arn",
					ContainerName:  "canary-container-name",
					ContainerPort:  80,
				},
			},
			expected: []*types.LoadBalancer{
				{
					TargetGroupArn: aws.String("primary-target-group-arn"),
					ContainerName:  aws.String("primary-container-name"),
					ContainerPort:  aws.Int32(80),
				},
				{
					TargetGroupArn: aws.String("canary-target-group-arn"),
					ContainerName:  aws.String("canary-container-name"),
					ContainerPort:  aws.Int32(80),
				},
			},
			expectedErr: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			primary, canary, err := loadTargetGroups(tc.cfg)
			assert.Equal(t, tc.expectedErr, err != nil)
			assert.Equal(t, tc.expected[0], primary)
			assert.Equal(t, tc.expected[1], canary)
		})
	}
}
