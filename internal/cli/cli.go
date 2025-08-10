package cli

import (
	"fmt"

	"github.com/arian-press2015/uniac/internal/plugins"
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
	var rootCmd = generateRootCmd()

	registerValidateCli(rootCmd)
	registerGenerateCli(rootCmd, c)
	registerPluginsCli(rootCmd)

	return rootCmd.Execute()
}
