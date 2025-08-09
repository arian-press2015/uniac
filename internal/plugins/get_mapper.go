package plugins

import (
	"fmt"

	"github.com/arian-press2015/uniac/pkg/plugins"
)

func GetMapper(pm *PluginManager, provider string, iac string) (*plugins.MapperPluginInterface, error) {
	metadata := plugins.MapperMetadata{Provider: provider, IaC: iac}

	plugin, err := pm.FindPlugin(
		plugins.PluginKindIaCMapper,
		metadata,
	)

	if err != nil {
		return nil, err
	}

	pluginInstance, ok := plugin.Instance.(*plugins.MapperPluginInterface)
	if !ok {
		return nil, fmt.Errorf("Plugin %s at %s has incompatible type, expected %T", plugin.Kind, plugin.Path, new(plugins.MapperPluginInterface))
	}

	return pluginInstance, nil
}
