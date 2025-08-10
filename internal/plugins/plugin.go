package plugins

import (
	"fmt"
	"os"

	"github.com/arian-press2015/uniac/pkg/plugins"
)

type PluginStatus string

const (
	PluginStatusSuccess PluginStatus = "success"
	PluginStatusFailure PluginStatus = "failure"
)

type Plugin struct {
	Kind     plugins.PluginKind
	Path     string
	Name     string
	Status   PluginStatus
	Metadata interface{}
	Instance interface{}
}

func (p *Plugin) String() string {
	metadataStr := "N/A"
	if p.Metadata != nil {
		metadataStr = fmt.Sprintf("%v", p.Metadata)
	}
	return fmt.Sprintf("Plugin{Kind: %s, Name: %s, Path: %s, Status: %s, Metadata: %s}",
		p.Kind, p.Name, p.Path, p.Status, metadataStr)
}

func (p *Plugin) Delete() error {
	err := os.Remove(p.Path)
	if err != nil {
		return fmt.Errorf("failed to delete the plugin %s: %v", p.Path, err)
	}
	return nil
}
