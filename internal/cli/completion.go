package cli

import "github.com/spf13/cobra"

func NewCompletionCmd() *cobra.Command {
	return newStubCmd("completion", "Generate shell completion scripts")
}
