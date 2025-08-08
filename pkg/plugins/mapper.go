package plugins

import "github.com/arian-press2015/uniac/pkg/core"

type MapperMetadata struct {
	IaC      string
	Provider string
}

type MapperPluginInterface interface {
	GetMetadata() MapperMetadata
	Generate(w *core.World) (string, error)
}
