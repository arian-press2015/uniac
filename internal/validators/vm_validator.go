package validators

import "fmt"

func (config *Config) validateVM() error {
	for i, vm := range config.VM {

		for _, disk := range vm.Disks {
			_, err := config.findDiskByName(disk)
			if err != nil {
				return fmt.Errorf("no disk found with the name of %s in vm %s at position %d", disk, vm.Name, i)
			}
		}

		_, err := config.findNetworkByName(vm.Network)
		if err != nil {
			return fmt.Errorf("no network found with the name of %s in vm %s at position %d", vm.Network, vm.Name, i)
		}
	}
	return nil
}
