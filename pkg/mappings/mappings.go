package mappings

import (
	"fmt"
	"os"
	"path/filepath"
	"plugin"
	"strings"

	"github.com/arian-press2015/uniac/pkg/core"
)

func GenerateIaCConfig(w *core.World, provider string, iac string) (string, error) {
	pluginsDir := os.Getenv("UNIAC_PLUGINS_DIR")
	if pluginsDir == "" {
		pluginsDir = filepath.Join(os.Getenv("HOME"), ".uniac", "plugins")
	}

	files, err := os.ReadDir(pluginsDir)
	if err != nil {
		return "", fmt.Errorf("failed to read plugins dir: %v", err)
	}

	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".so") {
			continue
		}

		p, err := plugin.Open(filepath.Join(pluginsDir, file.Name()))
		if err != nil {
			continue
		}

		symMapper, err := p.Lookup("Mapper")
		if err != nil {
			continue
		}

		mapperPtr, ok := symMapper.(*Mapper)
		if !ok {
			continue
		}
		mapper := *mapperPtr

		if mapper.Provider() == provider && mapper.IaC() == iac {
			return mapper.Generate(w)
		}
	}

	return "", fmt.Errorf("no mapper found for provider: %v, iac: %v", provider, iac)
}
