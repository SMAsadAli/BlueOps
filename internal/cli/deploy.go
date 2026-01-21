package cli

import "github.com/spf13/cobra"

func NewDeployCmd() *cobra.Command {
	return newStubCmd("deploy", "Deploy a service")
}
