package loader

import (
	"fmt"

	"github.com/arian-press2015/uniac/internal/parsers"
	"github.com/arian-press2015/uniac/internal/validators"
	"github.com/arian-press2015/uniac/pkg/core"
)

func Load(filepath string) (*core.World, error) {
	parser, err := parsers.NewParser(filepath)
	if err != nil {
		return nil, err
	}

	configMap, err := parser.Parse(filepath)
	if err != nil {
		return nil, err
	}

	config, err := validators.NewConfig(configMap)
	if err != nil {
		return nil, err
	}

	err = validators.Validate(config)
	if err != nil {
		return nil, err
	}

	w, err := core.NewWorld(config)
	if err != nil {
		return nil, fmt.Errorf("error creating World: %v", err)
	}

	return w, nil
}
