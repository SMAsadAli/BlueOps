package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

type LoadOptions struct {
	ConfigPath  string
	EnvOverride string
}

func Load(opts LoadOptions) (*Config, error) {
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetEnvPrefix("BLUEOPS")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	bindEnvKeys(v)
	v.AutomaticEnv()

	v.SetDefault("current_env", "dev")
	v.SetDefault("service_registry_path", "examples/services.yaml")

	if home, err := os.UserHomeDir(); err == nil {
		globalPath := filepath.Join(home, ".blueops", "config.yaml")
		if err := mergeConfigFile(v, globalPath); err != nil {
			return nil, err
		}
	}

	if opts.ConfigPath != "" {
		if err := mergeRequiredConfigFile(v, opts.ConfigPath); err != nil {
			return nil, err
		}
	} else {
		if err := mergeConfigFile(v, ".blueops.yaml"); err != nil {
			return nil, err
		}
	}

	if opts.EnvOverride != "" {
		v.Set("current_env", opts.EnvOverride)
	}

	envName := v.GetString("current_env")
	if envName != "" {
		envPath := fmt.Sprintf(".blueops.%s.yaml", envName)
		if err := mergeConfigFile(v, envPath); err != nil {
			return nil, err
		}
	}

	if opts.EnvOverride != "" {
		v.Set("current_env", opts.EnvOverride)
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("decode config: %w", err)
	}
	return &cfg, nil
}

func bindEnvKeys(v *viper.Viper) {
	keys := []string{
		"current_env",
		"service_registry_path",
		"control_plane_url",
		"auth.provider",
		"auth.profile",
	}
	for _, key := range keys {
		_ = v.BindEnv(key)
	}
}

func mergeConfigFile(v *viper.Viper, path string) error {
	f, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("open config %s: %w", path, err)
	}
	defer f.Close()

	if err := v.MergeConfig(f); err != nil {
		return fmt.Errorf("load config %s: %w", path, err)
	}
	return nil
}

func mergeRequiredConfigFile(v *viper.Viper, path string) error {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("config file not found: %s", path)
		}
		return err
	}
	return mergeConfigFile(v, path)
}
