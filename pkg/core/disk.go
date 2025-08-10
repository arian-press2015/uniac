package core

import "github.com/arian-press2015/uniac/internal/validators"

type Disk struct {
	Name       string
	Size       string
	Type       string
	Filesystem string
	Tags       map[string]string
}

func NewDisk(configDisk *validators.Disk) *Disk {
	return &Disk{
		Name:       configDisk.Name,
		Size:       configDisk.Size,
		Type:       configDisk.Type,
		Tags:       configDisk.Tags,
	}
}
