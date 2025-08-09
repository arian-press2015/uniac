package cli

import (
	"fmt"
	"os"

	"github.com/arian-press2015/uniac/internal/loader"
	"github.com/arian-press2015/uniac/internal/mappings"
	"github.com/arian-press2015/uniac/internal/plugins"
	"github.com/spf13/cobra"
)

type CLI struct {
	pm *plugins.PluginManager
}

func NewCLI() (*CLI, error) {
	pm := plugins.NewPluginManager()
	if err := pm.LoadPlugins(); err != nil {
		return nil, fmt.Errorf("failed to initialize plugins: %v", err)
	}
	return &CLI{pm: pm}, nil
}

func (c *CLI) RunCLI() error {
	var rootCmd = &cobra.Command{
		Use:   "uniac",
		Short: "uniac - Unified Infrastructure as Code",
		Long:  "uniac is a tool for managing infrastructure as code using a unified language",
	}

	var validateCmd = &cobra.Command{
		Use:   "validate",
		Short: "Validate a configuration file",
		Long:  "Validate reads a configuration file (default: infra.yaml) and shows errors if any found",
		Run: func(cmd *cobra.Command, args []string) {
			filepath, err := cmd.Flags().GetString("file")
			if err != nil {
				fmt.Println("Error getting file flag:", err)
				os.Exit(1)
			}

			_, err = loader.Load(filepath)
			if err != nil {
				fmt.Println("Error loading and validating:", err)
				os.Exit(1)
			}

			fmt.Println("World parsed and validated successfully")
		},
	}

	validateCmd.Flags().StringP("file", "f", "infra.yaml", "Path to the configuration file")

	var generateCmd = &cobra.Command{
		Use:   "generate <provider> <iac>",
		Short: "Generate infrastructure configuration",
		Long:  "Generate creates an infrastructure configuration for the specified provider and IaC tool using a configuration file",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			provider := args[0]
			iac := args[1]

			filepath, err := cmd.Flags().GetString("file")
			if err != nil {
				fmt.Println("Error getting file flag:", err)
				os.Exit(1)
			}

			w, err := loader.Load(filepath)
			if err != nil {
				fmt.Println("Error loading and validating:", err)
				os.Exit(1)
			}

			fmt.Println("World parsed and validated successfully")

			config, err := mappings.GenerateIaCConfig(c.pm, w, provider, iac)
			if err != nil {
				fmt.Println("Error in generating IaC output:", err)
				os.Exit(1)
			}

			fmt.Println("results:", config)
		},
	}

	generateCmd.Flags().StringP("file", "f", "infra.yaml", "Path to the configuration file")

	rootCmd.AddCommand(validateCmd)
	rootCmd.AddCommand(generateCmd)

	return rootCmd.Execute()
}
