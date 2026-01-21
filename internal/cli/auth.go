package cli

import "github.com/spf13/cobra"

func NewAuthCmd() *cobra.Command {
	cmd := newStubCmd("auth", "Manage authentication")
	cmd.AddCommand(
		NewAuthLoginCmd(),
		NewAuthStatusCmd(),
		NewAuthLogoutCmd(),
	)
	return cmd
}

func NewAuthLoginCmd() *cobra.Command {
	return newStubCmd("login", "Authenticate and store credentials")
}

func NewAuthStatusCmd() *cobra.Command {
	return newStubCmd("status", "Show authentication status")
}

func NewAuthLogoutCmd() *cobra.Command {
	return newStubCmd("logout", "Remove stored credentials")
}
