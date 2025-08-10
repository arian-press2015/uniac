package cli

import "github.com/spf13/cobra"

func generateRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "uniac",
		Short: "uniac - Unified Infrastructure as Code",
		Long:  "uniac is a tool for managing infrastructure as code using a unified language",
	}
}
