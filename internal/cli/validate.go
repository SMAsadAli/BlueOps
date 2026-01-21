package cli

import "github.com/spf13/cobra"

func NewValidateCmd() *cobra.Command {
	return newStubCmd("validate", "Validate configuration and guardrails")
}
