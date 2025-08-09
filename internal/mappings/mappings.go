package mappings

import (
	"fmt"
	"os"

	"github.com/arian-press2015/uniac/internal/plugins"
	"github.com/arian-press2015/uniac/pkg/core"
)

func GenerateIaCConfig(w *core.World, provider string, iac string) (string, error) {
	pm := plugins.NewPluginManager()
	if err := pm.LoadPlugins(); err != nil {
		fmt.Println("Failed to load plugins:", err)
		os.Exit(1)
	}

	mapper, err := plugins.GetMapper(pm, provider, iac)
	if err != nil {
		return "", err
	}

	iacConfig, err := (*mapper).Generate(w)
	if err != nil {
		return "", fmt.Errorf("no mapper found for provider: %v, iac: %v", provider, iac)
	}

	return iacConfig, nil
}
