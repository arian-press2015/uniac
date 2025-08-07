package cli

import (
	"fmt"
	"os"

	"github.com/arian-press2015/uniac/internal/loader"
	"github.com/arian-press2015/uniac/pkg/mappings"
	"github.com/spf13/cobra"
)

func RunCLI() error {
	var rootCmd = &cobra.Command{
		Use:   "uniac",
		Short: "uniac - Unified Infrastructure as Code",
		Long:  "uniac is a tool for managing infrastructure as code using a unified language",
	}

	var validateCmd = &cobra.Command{
		Use:   "validate [config-file]",
		Short: "Validate a configuration file",
		Long:  "Validate reads a configuration file(default: infra.yaml) and shows errors if any found",
		Run: func(cmd *cobra.Command, args []string) {
			var filepath string
			if len(args) > 0 {
				filepath = args[0]
			} else {
				filepath = "infra.yaml"
			}

			w, err := loader.Load(filepath)
			if err != nil {
				fmt.Println("Error loading and validating:", err)
				os.Exit(1)
			}

			fmt.Println("World parsed and validated successfully")

			config, err := mappings.GenerateIaCConfig(w, "example-cloud-provider", "terraform")

			fmt.Println("results:", config, err)
		},
	}

	rootCmd.AddCommand(validateCmd)

	return rootCmd.Execute()
}
