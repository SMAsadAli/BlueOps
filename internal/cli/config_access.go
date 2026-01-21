package cli

import (
	"errors"

	"github.com/SMAsadAli/blueops/internal/config"
	"github.com/spf13/cobra"
)

func getConfig(cmd *cobra.Command) (*config.Config, error) {
	cfg, ok := config.FromContext(cmd.Context())
	if !ok || cfg == nil {
		return nil, errors.New("config not loaded")
	}
	return cfg, nil
}
