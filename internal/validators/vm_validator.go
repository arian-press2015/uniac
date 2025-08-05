package validators

import (
	"fmt"

	"github.com/arian-press2015/uniac/pkg/core"
)

type VMValidator struct{}

func (vv *VMValidator) Validate(data interface{}, target *[]core.VM) error {
	vms, ok := data.([]interface{})
	if !ok {
		return fmt.Errorf("vms must be an array")
	}

	if *target == nil {
		*target = make([]core.VM, 0)
	}

	for i, vm := range vms {
		vmMap, ok := vm.(map[string]interface{})
		if !ok {
			return fmt.Errorf("vm at index %d must be a map", i)
		}

		name, _ := vmMap["name"].(string)
		cpu, _ := vmMap["cpu"].(int)
		ram, _ := vmMap["ram"].(int)

		if name == "" || cpu <= 0 || ram <= 0 {
			return fmt.Errorf("vm %d missing required field(s) 'name' or 'cpu' or 'ram'", i)
		}

		*target = append(*target, core.VM{
			Name: name,
			CPU:  cpu,
			RAM:  ram,
		})
	}

	return nil
}
