package validators

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
	validate.RegisterValidation("storage_size", func(fl validator.FieldLevel) bool {
		value := fl.Field().String()
		_, err := StorageSizeValidator(value)
		return err == nil
	})
}

func Validate(config *Config) error {
	if err := validate.Struct(config); err != nil {
		return fmt.Errorf("config validation failed: %v", err)
	}

	return nil
}
