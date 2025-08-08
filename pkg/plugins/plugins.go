package plugins

type PluginKind string

type PluginSymbol string

// list of supported plugin kinds
const (
	PluginKindIaCMapper PluginKind = "IaCMapper"
)

// the variable name that is read from the plugin
const (
	PluginSymbolIaCMapper PluginSymbol = "Mapper"
)

type PluginInterface interface {
	GetMetadata() interface{}
}