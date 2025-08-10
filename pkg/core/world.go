package core

import "github.com/arian-press2015/uniac/internal/validators"

type World struct {
	VM       []VM
	Disk     []Disk
	Database []Database
	Network   []Network
}

func NewWorld(config *validators.Config) (*World, error) {
	vms := make([]VM, len(config.VM))
	for i, vm := range config.VM {
		vms[i] = *NewVM(&vm)
	}

	disks := make([]Disk, len(config.Disk))
	for i, disk := range config.Disk {
		disks[i] = *NewDisk(&disk)
	}

	databases := make([]Database, len(config.Database))
	for i, db := range config.Database {
		databases[i] = *NewDatabase(&db)
	}

	networks := make([]Network, len(config.Network))
	for i, net := range config.Network {
		networks[i] = *NewNetwork(&net)
	}

	w := &World{
		VM:       vms,
		Disk:     disks,
		Database: databases,
		Network:   networks,
	}

	return w, nil
}
