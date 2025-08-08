package plugins

import "github.com/arian-press2015/uniac/pkg/plugins"

func GetMapper(provider string, iac string) (*plugins.MapperPluginInterface, error) {
	metadata := plugins.MapperMetadata{Provider: provider, IaC: iac}

	return FindPlugin[plugins.MapperPluginInterface](
		plugins.PluginKindIaCMapper,
		metadata,
	)
}
