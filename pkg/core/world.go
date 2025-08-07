package core

import "github.com/arian-press2015/uniac/internal/validators"

type World struct {
	VMs   []VM
	Disks []Disk
}

func NewWorld(config *validators.Config) (*World, error) {
	vms := make([]VM, len(config.VMs))
	for i, vm := range config.VMs {
		vms[i] = *NewVM(&vm)
	}

	disks := make([]Disk, len(config.Disks))
	for i, disk := range config.Disks {
		disks[i] = *NewDisk(&disk)
	}

	w := &World{
		VMs:   vms,
		Disks: disks,
	}

	return w, nil
}
