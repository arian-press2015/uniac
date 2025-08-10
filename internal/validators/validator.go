package validators

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func (config *Config) Validate() error {
	if err := config.validateStructure(); err != nil {
		return err
	}

	if err := config.validateSemantics(); err != nil {
		return err
	}

	return nil
}

// validates the config on a shallow level like struct structure and field types
func (config *Config) validateStructure() error {
	validate = validator.New()
	validate.RegisterValidation("storage_size", func(fl validator.FieldLevel) bool {
		value := fl.Field().String()
		_, err := StorageSizeValidator(value)
		return err == nil
	})

	if err := validate.Struct(config); err != nil {
		return fmt.Errorf("config validation failed: %v", err)
	}

	return nil
}

// validates the config deeply like logical relations and domain-specific validations
func (config *Config) validateSemantics() error {
	err := config.validateVM()
	if err != nil {
		return err
	}

	err = config.validateDisk()
	if err != nil {
		return err
	}

	err = config.validateDB()
	if err != nil {
		return err
	}

	err = config.validateNetwork()
	if err != nil {
		return err
	}

	return nil
}
