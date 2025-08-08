package plugins

import (
	"github.com/arian-press2015/uniac/pkg/plugins"
)

var PluginKindToSymbol = map[plugins.PluginKind]plugins.PluginSymbol{
	plugins.PluginKindIaCMapper: plugins.PluginSymbolIaCMapper,
}

var allKinds = []plugins.PluginKind{plugins.PluginKindIaCMapper}

var PluginRegistry = make(map[plugins.PluginKind][]*Plugin)
