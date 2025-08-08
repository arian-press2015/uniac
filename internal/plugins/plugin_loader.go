package plugins

import (
	"fmt"
	"os"
	"path/filepath"
	"plugin"
	"strings"
)

type VerifyFunc[T any] func(*T) bool

func LoadPlugin[T any](symbolName string, verify VerifyFunc[T]) (*T, error) {
	pluginsDir := os.Getenv("UNIAC_PLUGINS_DIR")
	if pluginsDir == "" {
		pluginsDir = filepath.Join(os.Getenv("HOME"), ".uniac", "plugins")
	}

	files, err := os.ReadDir(pluginsDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read plugins dir: %v", err)
	}

	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".so") {
			continue
		}

		p, err := plugin.Open(filepath.Join(pluginsDir, file.Name()))
		if err != nil {
			continue 
		}

		sym, err := p.Lookup(symbolName)
		if err != nil {
			continue
		}

		pluginVal, ok := sym.(*T)
		if !ok {
			continue
		}

		if verify != nil && !verify(pluginVal) {
			continue
		}

		return pluginVal, nil
	}

	return nil, fmt.Errorf("no plugin found with symbol %s of type %T", symbolName, new(T))
}
