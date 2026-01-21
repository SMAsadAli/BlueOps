package cli

import (
	"encoding/json"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func NewConfigCmd() *cobra.Command {
	cmd := newStubCmd("config", "Manage configuration")
	cmd.AddCommand(NewConfigViewCmd())
	return cmd
}

func NewConfigViewCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "view",
		Short: "Show the resolved configuration",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := getConfig(cmd)
			if err != nil {
				return err
			}
			jsonOut, _ := cmd.Flags().GetBool("json")
			if jsonOut {
				enc := json.NewEncoder(cmd.OutOrStdout())
				enc.SetIndent("", "  ")
				return enc.Encode(cfg)
			}
			data, err := yaml.Marshal(cfg)
			if err != nil {
				return err
			}
			_, err = cmd.OutOrStdout().Write(data)
			return err
		},
	}
}
