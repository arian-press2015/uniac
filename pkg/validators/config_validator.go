package validators

import (
	"fmt"

	"github.com/arian-press2015/uniac/pkg/core"
)

type ConfigValidator struct {
	VMValidator   *VMValidator
	DiskValidator *DiskValidator
}

func (cv *ConfigValidator) Validate(data interface{}, target *core.World) error {
	config, ok := data.(map[string]interface{})
	if !ok {
		return fmt.Errorf("config must be a map")
	}

	if vms, ok := config["vms"]; ok {
		if err := cv.VMValidator.Validate(vms, &target.VMs); err != nil {
			return fmt.Errorf("vms validation failed: %v", err)
		}
	}

	if disks, ok := config["disks"]; ok {
		if err := cv.DiskValidator.Validate(disks, &target.Disks); err != nil {
			return fmt.Errorf("disks validation failed: %v", err)
		}
	}

	return nil
}
