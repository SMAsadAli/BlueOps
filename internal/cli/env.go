package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

func NewEnvCmd() *cobra.Command {
	cmd := newStubCmd("env", "Manage environments")
	cmd.AddCommand(
		NewEnvListCmd(),
		NewEnvUseCmd(),
		NewEnvCurrentCmd(),
	)
	return cmd
}

func NewEnvListCmd() *cobra.Command {
	return newStubCmd("list", "List environments")
}

func NewEnvUseCmd() *cobra.Command {
	return newStubCmd("use", "Select an environment")
}

func NewEnvCurrentCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "current",
		Short: "Show the active environment",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := getConfig(cmd)
			if err != nil {
				return err
			}
			jsonOut, _ := cmd.Flags().GetBool("json")
			if jsonOut {
				payload := map[string]string{"current_env": cfg.CurrentEnv}
				enc := json.NewEncoder(cmd.OutOrStdout())
				enc.SetIndent("", "  ")
				return enc.Encode(payload)
			}
			_, err = fmt.Fprintln(cmd.OutOrStdout(), cfg.CurrentEnv)
			return err
		},
	}
}
