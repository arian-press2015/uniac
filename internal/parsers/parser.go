package parsers

import (
	"fmt"
)

type Parser interface {
	Parse(filepath string) (map[string]interface{}, error)
}

func NewParser(filePath string) (Parser, error) {
	if ext := filePath[len(filePath)-5:]; ext == ".yaml" || ext == ".yml" {
		return &YAMLParser{}, nil
	} else if ext := filePath[len(filePath)-5:]; ext == ".toml" {
		return &TOMLParser{}, nil
	}
	return nil, fmt.Errorf("unsupported file format: %s", filePath)
}
