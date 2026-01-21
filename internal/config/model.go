package config

type Config struct {
	CurrentEnv          string                 `mapstructure:"current_env" yaml:"current_env" json:"current_env"`
	Environments        map[string]Environment `mapstructure:"environments" yaml:"environments" json:"environments"`
	ServiceRegistryPath string                 `mapstructure:"service_registry_path" yaml:"service_registry_path" json:"service_registry_path"`
	ControlPlaneURL     string                 `mapstructure:"control_plane_url" yaml:"control_plane_url" json:"control_plane_url"`
	Auth                AuthConfig             `mapstructure:"auth" yaml:"auth" json:"auth"`
}

type Environment struct {
	Description     string `mapstructure:"description" yaml:"description" json:"description"`
	ControlPlaneURL string `mapstructure:"control_plane_url" yaml:"control_plane_url" json:"control_plane_url"`
}

type AuthConfig struct {
	Provider string `mapstructure:"provider" yaml:"provider" json:"provider"`
	Profile  string `mapstructure:"profile" yaml:"profile" json:"profile"`
}
