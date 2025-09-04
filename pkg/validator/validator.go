package validator

import "github.com/go-playground/validator"

type Validator struct {
	v *validator.Validate
}

func New() *Validator {
	return &Validator{
		v: validator.New(),
	}
}
