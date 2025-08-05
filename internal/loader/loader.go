package loader

import (
	"fmt"

	"github.com/arian-press2015/uniac/internal/parsers"
	"github.com/arian-press2015/uniac/internal/validators"
	"github.com/arian-press2015/uniac/pkg/core"
)

type Loader struct {
	Parser          parsers.Parser
	ConfigValidator *validators.ConfigValidator
}

func NewLoader(filepath string) (*Loader, error) {
	parser, err := parsers.NewParser(filepath)

	if err != nil {
		return nil, err
	}

	validator := &validators.ConfigValidator{
		VMValidator:   &validators.VMValidator{},
		DiskValidator: &validators.DiskValidator{},
	}

	return &Loader{Parser: parser, ConfigValidator: validator}, nil
}

func (l *Loader) Load(filepath string) (*core.World, error) {
	config, err := l.Parser.Parse(filepath)
	if err != nil {
		return nil, err
	}

	w, err := core.NewWorld(config)
	if err != nil {
		return nil, fmt.Errorf("error creating World: %v", err)
	}

	if err := l.ConfigValidator.Validate(config, w); err != nil {
		return nil, err
	}

	return w, nil
}
