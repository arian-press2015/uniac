package plugins

import (
	"fmt"
	"log"
	"reflect"

	"github.com/arian-press2015/uniac/pkg/plugins"
)

func FindPlugin[T any](kind plugins.PluginKind, desiredMeta interface{}) (*T, error) {
	plugins, ok := PluginRegistry[kind]
	if !ok || len(plugins) == 0 {
		return nil, fmt.Errorf("no plugins found for kind %s", kind)
	}

	for _, plugin := range plugins {
		if plugin.Status != PluginStatusSuccess {
			continue
		}

		if plugin.Instance == nil {
			log.Printf("No instance for plugin %s at %s", plugin.Kind, plugin.Path)
			continue
		}
		pluginVal, ok := plugin.Instance.(*T)
		if !ok {
			log.Printf("Plugin %s at %s has incompatible type, expected %T", plugin.Kind, plugin.Path, new(T))
			continue
		}

		if reflect.TypeOf(plugin.Metadata) == reflect.TypeOf(desiredMeta) {
			if reflect.DeepEqual(plugin.Metadata, desiredMeta) {
				return pluginVal, nil
			}
		}
	}

	return nil, fmt.Errorf("no plugin found with kind %s and matching metadata", kind)
}
