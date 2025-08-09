package mappings

import (
	"fmt"

	"github.com/arian-press2015/uniac/internal/plugins"
	"github.com/arian-press2015/uniac/pkg/core"
)

func GenerateIaCConfig(pm *plugins.PluginManager, w *core.World, provider string, iac string) (string, error) {
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
