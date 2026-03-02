package deployment

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"

	ecsconfig "github.com/pipe-cd/pipecd/pkg/app/pipedv1/plugin/ecs/config"
	sdk "github.com/pipe-cd/piped-plugin-sdk-go"
	"go.uber.org/zap"
)

func TestECSSyncStage(t *testing.T) {
	t.Parallel()

	var (
		clusterARN   = os.Getenv("ECS_CLUSTER_ARN")
		subnetIDsStr = os.Getenv("SUBNET_IDS")
		sgIDsStr     = os.Getenv("SECURITY_GROUP_IDS")

		subnetIDs []string
		sgIDs     []string
	)

	if clusterARN == "" {
		t.Skip("ECS_CLUSTER_ARN is not set, skipping ECS sync stage test")
	}
	if subnetIDsStr == "" {
		t.Skip("SUBNET_IDS is not set, skipping ECS sync stage test")
	}
	for _, id := range strings.Split(subnetIDsStr, ",") {
		subnetIDs = append(subnetIDs, strings.TrimSpace(id))
	}
	if len(subnetIDs) == 0 {
		t.Skip("SUBNET_IDS is empty, skipping ECS sync stage test")
	}
	if sgIDsStr == "" {
		t.Skip("SECURITY_GROUP_IDS is not set, skipping ECS sync stage test")
	}
	for _, id := range strings.Split(sgIDsStr, ",") {
		sgIDs = append(sgIDs, strings.TrimSpace(id))
	}
	if len(sgIDs) == 0 {
		t.Skip("SECURITY_GROUP_IDS is empty, skipping ECS sync stage test")
	}

	logger, err := zap.NewDevelopment()
	if err != nil {
		t.Fatalf("Failed to create logger: %v", err)
	}

	testcases := []struct {
		name  string
		input *sdk.ExecuteStageInput[ecsconfig.ECSApplicationSpec]
		dts   []*sdk.DeployTarget[ecsconfig.ECSDeployTargetConfig]
	}{
		{
			name: "standalone task",
			input: &sdk.ExecuteStageInput[ecsconfig.ECSApplicationSpec]{
				Logger: logger,
				Client: sdk.NewClient(nil, "", "", "", &mockStageLogPersister{logger: logger}, nil),
				Request: sdk.ExecuteStageRequest[ecsconfig.ECSApplicationSpec]{
					TargetDeploymentSource: sdk.DeploymentSource[ecsconfig.ECSApplicationSpec]{
						ApplicationDirectory: "testdata/standalone-task",
						ApplicationConfig: &sdk.ApplicationConfig[ecsconfig.ECSApplicationSpec]{
							Spec: &ecsconfig.ECSApplicationSpec{
								Input: ecsconfig.ECSDeploymentInput{
									TaskDefinitionFile:    "taskdef.yaml",
									ServiceDefinitionFile: "",
									RunStandaloneTask:     true,
									ClusterARN:            clusterARN,
									LaunchType:            "FARGATE",
									AwsVpcConfiguration: ecsconfig.ECSVpcConfiguration{
										Subnets:        subnetIDs,
										SecurityGroups: sgIDs,
									},
								},
							},
						},
					},
				},
			},
			dts: []*sdk.DeployTarget[ecsconfig.ECSDeployTargetConfig]{
				{
					Name: "test-deploy-target",
					// Use default profile for test
					Config: ecsconfig.ECSDeployTargetConfig{
						Region: "us-east-1",
					},
				},
			},
		},
		{
			name: "service without load balancer",
			input: &sdk.ExecuteStageInput[ecsconfig.ECSApplicationSpec]{
				Logger: logger,
				Client: sdk.NewClient(nil, "", "", "", &mockStageLogPersister{logger: logger}, nil),
				Request: sdk.ExecuteStageRequest[ecsconfig.ECSApplicationSpec]{
					TargetDeploymentSource: sdk.DeploymentSource[ecsconfig.ECSApplicationSpec]{
						ApplicationDirectory: "testdata/service-without-lb",
						ApplicationConfig: &sdk.ApplicationConfig[ecsconfig.ECSApplicationSpec]{
							Spec: &ecsconfig.ECSApplicationSpec{
								Input: ecsconfig.ECSDeploymentInput{
									TaskDefinitionFile:    "taskdef.yaml",
									ServiceDefinitionFile: "servicedef.yaml",
									ClusterARN:            clusterARN,
									LaunchType:            "FARGATE",
									AwsVpcConfiguration: ecsconfig.ECSVpcConfiguration{
										Subnets:        subnetIDs,
										SecurityGroups: sgIDs,
									},
									AccessType: "SERVICE_DISCOVERY",
								},
							},
						},
					},
				},
			},
			dts: []*sdk.DeployTarget[ecsconfig.ECSDeployTargetConfig]{
				{
					Name: "test-deploy-target",
					// Use default profile for test
					Config: ecsconfig.ECSDeployTargetConfig{
						Region: "us-east-1",
					},
				},
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			plugin := &ECSPlugin{}
			status := plugin.executeECSSyncStage(context.Background(), tc.input, tc.dts)
			if status != sdk.StageStatusSuccess {
				t.Errorf("Expected stage status to be success, got %s", status)
			}
		})
	}
}

type mockStageLogPersister struct {
	logger *zap.Logger
}

func (m *mockStageLogPersister) Write(b []byte) (int, error) {
	m.logger.Info(string(b))
	return len(b), nil
}

func (m *mockStageLogPersister) Info(msg string) {
	m.logger.Info(msg)
}

func (m *mockStageLogPersister) Infof(format string, args ...interface{}) {
	m.logger.Info(fmt.Sprintf(format, args...))
}

func (m *mockStageLogPersister) Error(msg string) {
	m.logger.Error(msg)
}

func (m *mockStageLogPersister) Errorf(format string, args ...interface{}) {
	m.logger.Error(fmt.Sprintf(format, args...))
}

func (m *mockStageLogPersister) Success(msg string) {
	m.logger.Info(msg)
}

func (m *mockStageLogPersister) Successf(format string, args ...interface{}) {
	m.logger.Info(fmt.Sprintf(format, args...))
}
