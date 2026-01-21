package cli

import "github.com/spf13/cobra"

func NewInitCmd() *cobra.Command {
	return newStubCmd("init", "Initialize BlueOps configuration")
}
