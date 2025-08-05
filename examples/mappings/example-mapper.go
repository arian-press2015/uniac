package main

import (
	"fmt"

	"github.com/arian-press2015/uniac/pkg/core"
	"github.com/arian-press2015/uniac/pkg/mappings"
)

type ExampleMapper struct{}

func (em *ExampleMapper) Provider() string {
	return "example-cloud-provider"
}

func (em *ExampleMapper) IaC() string {
	return "terraform"
}

func (em *ExampleMapper) Generate(w *core.World) (string, error) {
	hcl := "resource \"aws_instance\" \"example\" {\n"
	for _, vm := range w.VMs {
		hcl += fmt.Sprintf("  instance_type = \"t2.micro\"\n  ami = \"ami-123456\"\n  tags = { Name = \"%s\" }\n", vm.Name)
	}
	hcl += "}\n"
	return hcl, nil
}

var Mapper mappings.Mapper = &ExampleMapper{}

func main() {}
