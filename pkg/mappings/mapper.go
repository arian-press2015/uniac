package mappings

import "github.com/arian-press2015/uniac/pkg/core"

type Mapper interface{
	Provider() string
	IaC() string
	Generate(w *core.World) (string, error)
}