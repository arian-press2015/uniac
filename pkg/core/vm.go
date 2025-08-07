package core

import "github.com/arian-press2015/uniac/internal/validators"

type VM struct {
	Name   string
	Size   string
	Type   string
	Region string
	Image  string
	Disks  []string
	Tags   map[string]string
}

func NewVM(configVM *validators.VM) *VM {
	return &VM{
		Name:   configVM.Name,
		Size:   configVM.Size,
		Type:   configVM.Type,
		Region: configVM.Region,
		Image:  configVM.Image,
		Disks:  configVM.Disks,
		Tags:   configVM.Tags,
	}
}
