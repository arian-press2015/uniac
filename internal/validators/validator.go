package validators

type Validator interface {
	Validate(data interface{}, target *interface{}) error
}