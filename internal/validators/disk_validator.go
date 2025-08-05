package validators

import (
	"fmt"

	"github.com/arian-press2015/uniac/pkg/core"
)

type DiskValidator struct{}

func (dv *DiskValidator) Validate(data interface{}, target *[]core.Disk) error {
	disks, ok := data.([]interface{})
	if !ok {
		return fmt.Errorf("disks must be an array")
	}

	for i, disk := range disks {
		diskMap, ok := disk.(map[string]interface{})
		if !ok {
			return fmt.Errorf("disk at index %d must be a map", i)
		}

		if *target == nil {
			*target = make([]core.Disk, 0)
		}

		kind, _ := diskMap["kind"].(string)
		capacity, _ := diskMap["capacity"].(int)
		fs, _ := diskMap["fs"].(string)

		if kind == "" || capacity <= 0 || fs == "" {
			return fmt.Errorf("disk %d missing required field(s) 'kind' or 'capacity' or 'fs'", i)
		}

		*target = append(*target, core.Disk{
			Kind:     kind,
			Capacity: capacity,
			FS:       fs,
		})
	}

	return nil
}
