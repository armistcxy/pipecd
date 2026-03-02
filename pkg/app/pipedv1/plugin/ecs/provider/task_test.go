package provider

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/go-playground/assert/v2"
)

func TestParseTaskDefinition(t *testing.T) {
	t.Parallel()

	testcases := []struct {
		name        string
		input       string
		expected    types.TaskDefinition
		expectedErr bool
	}{
		{
			name: "yaml format input",
			input: `
family: nginx-canary-fam-1
compatibilities:
  - FARGATE
networkMode: awsvpc
memory: 512
cpu: 256
`,
			expected: types.TaskDefinition{
				Family:          aws.String("nginx-canary-fam-1"),
				Compatibilities: []types.Compatibility{types.CompatibilityFargate},
				NetworkMode:     types.NetworkModeAwsvpc,
				Memory:          aws.String("512"),
				Cpu:             aws.String("256"),
			},
		},
		{
			name: "json format input",
			input: `
{
  "family": "nginx-canary-fam-1",
  "compatibilities": [
    "FARGATE"
  ],
  "networkMode": "awsvpc",
  "memory": 512,
  "cpu": 256
}
`,
			expected: types.TaskDefinition{
				Family:          aws.String("nginx-canary-fam-1"),
				Compatibilities: []types.Compatibility{types.CompatibilityFargate},
				NetworkMode:     types.NetworkModeAwsvpc,
				Memory:          aws.String("512"),
				Cpu:             aws.String("256"),
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := parseTaskDefinition([]byte(tc.input))
			assert.Equal(t, tc.expectedErr, err != nil)
			assert.Equal(t, tc.expected, got)
		})
	}
}
