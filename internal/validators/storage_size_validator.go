package validators

import (
	"fmt"
	"regexp"
	"strings"
)

func StorageSizeValidator(sizeStr string) (uint64, error) {
	re := regexp.MustCompile(`^(\d+)(Gi|Mi|Ti)?$`)
	if !re.MatchString(sizeStr) {
		return 0, fmt.Errorf("invalid size format, expected [number][Gi|Mi|Ti], got %s", sizeStr)
	}

	var value float64
	var unit string
	_, err := fmt.Sscanf(sizeStr, "%f%s", &value, &unit)
	if err != nil {
		return 0, fmt.Errorf("failed to parse size: %v", err)
	}

	multiplier := uint64(1)
	unit = strings.ToLower(unit)
	switch unit {
	case "gi":
		multiplier = 1024 * 1024 * 1024
	case "mi":
		multiplier = 1024 * 1024
	case "ti":
		multiplier = 1024 * 1024 * 1024 * 1024
	case "":
		multiplier = 1024 * 1024 * 1024 // Default to GiB
	default:
		return 0, fmt.Errorf("unsupported unit: %s", unit)
	}

	if value <= 0 {
		return 0, fmt.Errorf("size must be positive, got %s", sizeStr)
	}

	return uint64(value * float64(multiplier)), nil
}
