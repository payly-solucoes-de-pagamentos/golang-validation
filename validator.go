package validation

type IValidator interface {
	Validate(instance interface{}) error
}
