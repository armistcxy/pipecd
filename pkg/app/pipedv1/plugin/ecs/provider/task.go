package provider

import (
	"os"

	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"sigs.k8s.io/yaml"
)

func loadTaskDefinition(path string) (types.TaskDefinition, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return types.TaskDefinition{}, err
	}
	return parseTaskDefinition(data)
}

func parseTaskDefinition(data []byte) (types.TaskDefinition, error) {
	var obj types.TaskDefinition
	if err := yaml.Unmarshal(data, &obj); err != nil {
		return types.TaskDefinition{}, err
	}
	return obj, nil
}
