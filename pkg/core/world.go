package core

import "github.com/arian-press2015/uniac/internal/validators"

type World struct {
	VMs       []VM
	Disks     []Disk
	Databases []Database
	Network   []Network
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

	databases := make([]Database, len(config.Databases))
	for i, db := range config.Databases {
		databases[i] = *NewDatabase(&db)
	}

	networks := make([]Network, len(config.Network))
	for i, net := range config.Network {
		networks[i] = *NewNetwork(&net)
	}

	w := &World{
		VMs:       vms,
		Disks:     disks,
		Databases: databases,
		Network:   networks,
	}

	return w, nil
}
