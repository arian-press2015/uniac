package plugins

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"plugin"
	"strings"

	"github.com/arian-press2015/uniac/pkg/plugins"
)

func LoadPlugins() error {
	pluginsDir := os.Getenv("UNIAC_PLUGINS_DIR")
	if pluginsDir == "" {
		pluginsDir = filepath.Join(os.Getenv("HOME"), ".uniac", "plugins")
	}

	for k := range PluginRegistry {
		delete(PluginRegistry, k)
	}

	files, err := os.ReadDir(pluginsDir)
	if err != nil {
		return fmt.Errorf("failed to read plugins dir: %v", err)
	}

	for _, file := range files {
		filePath := filepath.Join(pluginsDir, file.Name())

		if file.IsDir() || !strings.HasSuffix(file.Name(), ".so") {
			continue
		}

		// Extract kind from filename (e.g., "-IaCMapper.so")
		kindStr := ""
		for _, kind := range allKinds {
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
				Path:   filePath,
				Status: PluginStatusFailure,
			}
			PluginRegistry[plugins.PluginKind(kindStr)] = append(PluginRegistry[plugins.PluginKind(kindStr)], plugin)
			continue
		}

		symbol, ok := PluginKindToSymbol[plugins.PluginKind(kindStr)]
		if !ok {
			log.Printf("No symbol defined for kind %s in %s", kindStr, filePath)
			continue
		}

		sym, err := p.Lookup(string(symbol))
		if err != nil {
			log.Printf("Symbol %s not found in %s: %v", symbol, filePath, err)
			plugin := &Plugin{
				Kind:   plugins.PluginKind(kindStr),
				Path:   filePath,
				Status: PluginStatusFailure,
			}
			PluginRegistry[plugins.PluginKind(kindStr)] = append(PluginRegistry[plugins.PluginKind(kindStr)], plugin)
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
			Path:     filePath,
			Status:   PluginStatusSuccess,
			Metadata: metadata,
			Instance: sym,
		}
		PluginRegistry[plugins.PluginKind(kindStr)] = append(PluginRegistry[plugins.PluginKind(kindStr)], plugin)
	}

	return nil
}
