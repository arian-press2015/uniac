package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func registerPluginsCli(rootCmd *cobra.Command) {
	var pluginsCmd = &cobra.Command{
		Use:   "plugins",
		Short: "Manage plugins",
	}

	var listPluginsCmd = &cobra.Command{
		Use:   "list",
		Short: "List all loaded plugins",
		Long:  "List displays a table of all loaded plugins with their kinds, paths, statuses, and metadata",
		Run: func(cmd *cobra.Command, args []string) {
			c, err := NewCLI()
			if err != nil {
				fmt.Println("Error initializing CLI:", err)
				os.Exit(1)
			}
			fmt.Print(c.pm.String())
		},
	}

	var deletePluginCmd = &cobra.Command{
		Use:   "delete <plugin-name>",
		Short: "Delete a plugin by its name",
		Long:  "Delete removes a plugin from the registry using its file name",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			c, err := NewCLI()
			if err != nil {
				fmt.Println("Error initializing CLI:", err)
				os.Exit(1)
			}
			pluginName := args[0]

			if err := c.pm.DeletePlugin(pluginName); err != nil {
				fmt.Printf("Error deleting plugin: %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("Plugin %s deleted successfully\n", pluginName)
		},
	}

	pluginsCmd.AddCommand(listPluginsCmd)
	pluginsCmd.AddCommand(deletePluginCmd)

	rootCmd.AddCommand(pluginsCmd)
}
