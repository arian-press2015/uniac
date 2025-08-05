package main

import (
	"fmt"
	"os"

	"github.com/arian-press2015/uniac/internal/loader"
	"github.com/spf13/cobra"
)

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

		loader, err := loader.NewLoader(filepath)
		if err != nil {
			fmt.Println("Error creating config loader:", err)
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

func init() {
	rootCmd.AddCommand(validateCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
