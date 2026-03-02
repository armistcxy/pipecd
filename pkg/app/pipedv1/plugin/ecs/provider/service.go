package provider

import (
	"os"

	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"sigs.k8s.io/yaml"
)

func loadServiceDefinition(path string) (types.Service, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return types.Service{}, err
	}
	return parseServiceDefinition(data)
}

func parseServiceDefinition(data []byte) (types.Service, error) {
	var obj types.Service
	if err := yaml.Unmarshal(data, &obj); err != nil {
		return types.Service{}, err
	}

	if obj.ClusterArn == nil {
		// Rename cluster field to clusterArn if exist
		clusterArn, err := parseServiceDefinitionForCluster(data)
		if err != nil {
			return types.Service{}, err
		}
		obj.ClusterArn = &clusterArn
	}

	if obj.RoleArn == nil {
		// Rename role field to roleArn if exist
		roleArn, err := parseServiceDefinitionForRole(data)
		if err != nil {
			return types.Service{}, err
		}
		obj.RoleArn = &roleArn
	}

	return obj, nil
}

func parseServiceDefinitionForCluster(data []byte) (string, error) {
	var obj struct {
		Cluster string `json:"cluster"`
	}
	if err := yaml.Unmarshal(data, &obj); err != nil {
		return "", err
	}
	return obj.Cluster, nil
}

func parseServiceDefinitionForRole(data []byte) (string, error) {
	var obj struct {
		Role string `json:"role"`
	}
	if err := yaml.Unmarshal(data, &obj); err != nil {
		return "", err
	}
	return obj.Role, nil
}
