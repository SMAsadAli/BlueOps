package cli

import "github.com/spf13/cobra"

func NewServiceCmd() *cobra.Command {
	cmd := newStubCmd("service", "Manage services")
	cmd.AddCommand(
		NewServiceListCmd(),
		NewServiceDescribeCmd(),
	)
	return cmd
}

func NewServiceListCmd() *cobra.Command {
	return newStubCmd("list", "List services")
}

func NewServiceDescribeCmd() *cobra.Command {
	return newStubCmd("describe", "Describe a service")
}
