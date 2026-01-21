package cli

import (
	"github.com/SMAsadAli/blueops/internal/config"
	"github.com/spf13/cobra"
)

type RootOptions struct {
	ConfigPath string
	Env        string
	JSON       bool
	Verbose    bool
}

func NewRootCmd() *cobra.Command {
	opts := &RootOptions{}
	cmd := &cobra.Command{
		Use:          "blueops",
		Short:        "Enterprise-grade internal CLI for ops workflows",
		SilenceUsage: true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := config.Load(config.LoadOptions{
				ConfigPath:  opts.ConfigPath,
				EnvOverride: opts.Env,
			})
			if err != nil {
				return err
			}
			cmd.SetContext(config.WithConfig(cmd.Context(), cfg))
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	cmd.PersistentFlags().StringVar(&opts.ConfigPath, "config", "", "Path to config file")
	cmd.PersistentFlags().StringVar(&opts.Env, "env", "", "Environment name")
	cmd.PersistentFlags().BoolVar(&opts.JSON, "json", false, "Output JSON instead of table output")
	cmd.PersistentFlags().BoolVar(&opts.Verbose, "verbose", false, "Enable verbose logging")

	cmd.AddCommand(
		NewVersionCmd(),
		NewInitCmd(),
		NewAuthCmd(),
		NewEnvCmd(),
		NewServiceCmd(),
		NewDeployCmd(),
		NewValidateCmd(),
		NewConfigCmd(),
		NewCompletionCmd(),
	)

	return cmd
}
