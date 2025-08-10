package validators

import "fmt"

func (c *Config) findVMByName(name string) (*VM, error) {
	for i := range c.VM {
		if c.VM[i].Name == name {
			return &c.VM[i], nil
		}
	}
	return nil, fmt.Errorf("no VM found with name %s", name)
}

func (c *Config) findDiskByName(name string) (*Disk, error) {
	for i := range c.Disk {
		if c.Disk[i].Name == name {
			return &c.Disk[i], nil
		}
	}
	return nil, fmt.Errorf("no Disk found with name %s", name)
}

func (c *Config) findDBByName(name string) (*Database, error) {
	for i := range c.Database {
		if c.Database[i].Name == name {
			return &c.Database[i], nil
		}
	}
	return nil, fmt.Errorf("no Database found with name %s", name)
}

func (c *Config) findNetworkByName(name string) (*Network, error) {
	for i := range c.Network {
		if c.Network[i].Name == name {
			return &c.Network[i], nil
		}
	}
	return nil, fmt.Errorf("no Network found with name %s", name)
}
