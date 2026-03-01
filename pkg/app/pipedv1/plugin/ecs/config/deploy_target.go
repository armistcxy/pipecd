package config

// ECSDeployTargetConfig represents the per-deploy-target AWS configuration for ECS deployments
type ECSDeployTargetConfig struct {
	// AWS region where the ECS cluster is located
	// e.g., "us-west-2"
	Region string `json:"region"`

	// AWS profile to use from the credentials file
	// If empty, uses the default profile or "default" if AWS_PROFILE env var is not set
	Profile string `json:"profile,omitempty"`

	// Path to the AWS shared credentials file
	// e.g., "~/.aws/credentials"
	// If empty, uses the default location
	CredentialsFile string `json:"credentialsFile,omitempty"`

	// IAM role ARN to assume when accessing AWS resources
	// e.g., "arn:aws:iam::123456789:role/ecs-deployment-role"
	// Required when assuming a role across accounts
	RoleARN string `json:"roleARN,omitempty"`

	// Path to the OIDC token file for web identity federation
	// Required when RoleARN is set for OIDC-based authentication
	TokenFile string `json:"tokenFile,omitempty"`
}
