package core

type World struct {
	VMs   []VM
	Disks []Disk
}

func NewWorld(config map[string]interface{}) (*World, error) {
	w := &World{}

	return w, nil
}
