package cli

import "github.com/spf13/cobra"

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
	return newStubCmd("current", "Show the active environment")
}
