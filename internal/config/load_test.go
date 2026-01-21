package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadEnvOverridesConfig(t *testing.T) {
	tmp := t.TempDir()
	t.Setenv("HOME", tmp)

	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("getwd: %v", err)
	}
	if err := os.Chdir(tmp); err != nil {
		t.Fatalf("chdir: %v", err)
	}
	t.Cleanup(func() { _ = os.Chdir(cwd) })

	configPath := filepath.Join(tmp, ".blueops.yaml")
	configData := []byte("current_env: dev\nauth:\n  provider: file\n")
	if err := os.WriteFile(configPath, configData, 0o600); err != nil {
		t.Fatalf("write config: %v", err)
	}

	t.Setenv("BLUEOPS_CURRENT_ENV", "staging")
	t.Setenv("BLUEOPS_AUTH_PROVIDER", "keyring")

	cfg, err := Load(LoadOptions{})
	if err != nil {
		t.Fatalf("load: %v", err)
	}
	if cfg.CurrentEnv != "staging" {
		t.Fatalf("expected current_env staging, got %q", cfg.CurrentEnv)
	}
	if cfg.Auth.Provider != "keyring" {
		t.Fatalf("expected auth.provider keyring, got %q", cfg.Auth.Provider)
	}
}
