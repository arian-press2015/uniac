package plugins

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"plugin"
	"reflect"
	"strings"

	"github.com/arian-press2015/uniac/pkg/plugins"
	"github.com/olekukonko/tablewriter"
)

type PluginManager struct {
	pluginKindToSymbol map[plugins.PluginKind]plugins.PluginSymbol
	allKinds           []plugins.PluginKind
	pluginRegistry     map[plugins.PluginKind][]*Plugin
}

func NewPluginManager() *PluginManager {
	return &PluginManager{
		pluginKindToSymbol: map[plugins.PluginKind]plugins.PluginSymbol{
			plugins.PluginKindIaCMapper: plugins.PluginSymbolIaCMapper,
		},
		allKinds: []plugins.PluginKind{
			plugins.PluginKindIaCMapper,
		},
		pluginRegistry: make(map[plugins.PluginKind][]*Plugin),
	}
}

func (pm *PluginManager) String() string {
	var builder strings.Builder

	table := tablewriter.NewWriter(&builder)
	table.Header([]string{"Kind", "Name", "Path", "Status", "Metadata"})

	for kind, plugins := range pm.pluginRegistry {
		for _, plugin := range plugins {
			metadataStr := "N/A"
			if plugin.Metadata != nil {
				metadataStr = fmt.Sprintf("%v", plugin.Metadata)
			}

			table.Append([]string{
				string(kind),
				plugin.Name,
				plugin.Path,
				string(plugin.Status),
				metadataStr,
			})
		}
	}

	table.Render()

	return builder.String()
}

func (pm *PluginManager) DeletePlugin(pluginName string) error {
	for kind, plugins := range pm.pluginRegistry {
		for i, plugin := range plugins {
			if plugin.Name == pluginName {
				plugin.Delete()
				pm.pluginRegistry[kind] = append(plugins[:i], plugins[i+1:]...)
				if len(pm.pluginRegistry[kind]) == 0 {
					delete(pm.pluginRegistry, kind)
				}
				return nil
			}
		}
	}
	return fmt.Errorf("plugin with name %s not found", pluginName)
}

func (pm *PluginManager) LoadPlugins() error {
	pluginsDir := os.Getenv("UNIAC_PLUGINS_DIR")
	if pluginsDir == "" {
		pluginsDir = filepath.Join(os.Getenv("HOME"), ".uniac", "plugins")
	}

	// Clear existing registry
	for k := range pm.pluginRegistry {
		delete(pm.pluginRegistry, k)
	}

	// Read plugin directory
	files, err := os.ReadDir(pluginsDir)
	if err != nil {
		return fmt.Errorf("failed to read plugins dir: %v", err)
	}

	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".so") {
			continue
		}

		kindStr := ""
		filePath := filepath.Join(pluginsDir, file.Name())
		for _, kind := range pm.allKinds {
			if strings.HasSuffix(file.Name(), fmt.Sprintf("-%s.so", kind)) {
				kindStr = string(kind)
				break
			}
		}
		if kindStr == "" {
			log.Printf("Skipping %s: no recognized kind", file.Name())
			continue
		}

		p, err := plugin.Open(filePath)
		if err != nil {
			log.Printf("Failed to open plugin %s: %v", filePath, err)
			plugin := &Plugin{
				Kind:   plugins.PluginKind(kindStr),
				Name:   strings.TrimSuffix(file.Name(), ".so"),
				Path:   filePath,
				Status: PluginStatusFailure,
			}
			pm.pluginRegistry[plugins.PluginKind(kindStr)] = append(pm.pluginRegistry[plugins.PluginKind(kindStr)], plugin)
			continue
		}

		symbol, ok := pm.pluginKindToSymbol[plugins.PluginKind(kindStr)]
		if !ok {
			log.Printf("No symbol defined for kind %s in %s", kindStr, filePath)
			continue
		}

		sym, err := p.Lookup(string(symbol))
		if err != nil {
			log.Printf("Symbol %s not found in %s: %v", symbol, filePath, err)
			plugin := &Plugin{
				Kind:   plugins.PluginKind(kindStr),
				Name:   strings.TrimSuffix(file.Name(), ".so"),
				Path:   filePath,
				Status: PluginStatusFailure,
			}
			pm.pluginRegistry[plugins.PluginKind(kindStr)] = append(pm.pluginRegistry[plugins.PluginKind(kindStr)], plugin)
			continue
		}

		var metadata interface{}
		if getter, ok := sym.(*plugins.MapperPluginInterface); ok {
			metadata = (*getter).GetMetadata()
		} else {
			log.Printf("Plugin %s does not implement PluginInterface, metadata will be nil", filePath)
		}

		plugin := &Plugin{
			Kind:     plugins.PluginKind(kindStr),
			Name:     strings.TrimSuffix(file.Name(), ".so"),
			Path:     filePath,
			Status:   PluginStatusSuccess,
			Metadata: metadata,
			Instance: sym,
		}
		pm.pluginRegistry[plugins.PluginKind(kindStr)] = append(pm.pluginRegistry[plugins.PluginKind(kindStr)], plugin)
	}

	return nil
}

func (pm *PluginManager) FindPlugin(kind plugins.PluginKind, desiredMeta interface{}) (*Plugin, error) {
	plugins, ok := pm.pluginRegistry[kind]
	if !ok || len(plugins) == 0 {
		return nil, fmt.Errorf("no plugins found for kind %s", kind)
	}

	for _, plugin := range plugins {
		if plugin.Status == PluginStatusSuccess {
			if desiredMeta == nil || reflect.DeepEqual(plugin.Metadata, desiredMeta) {
				return plugin, nil
			}
		}
	}

	return nil, fmt.Errorf("no plugin found with kind %s and matching metadata", kind)
}
