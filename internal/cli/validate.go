package cli

import (
	"fmt"
	"os"

	"github.com/arian-press2015/uniac/internal/loader"
	"github.com/spf13/cobra"
)

func registerValidateCli(rootCmd *cobra.Command) {
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

	rootCmd.AddCommand(validateCmd)
}
