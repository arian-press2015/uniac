package parsers

import (
	"fmt"
	"os"

	"github.com/pelletier/go-toml/v2"
)

type TOMLParser struct{}

func (p *TOMLParser) Parse(filepath string) (map[string]interface{}, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	var parsed map[string]interface{}
	if err := toml.Unmarshal(data, &parsed); err != nil {
		return nil, err
	}
	return parsed, nil
}
