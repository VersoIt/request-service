package validator

func (v *Validator) Validate(i interface{}) error {
	if err := v.v.Struct(i); err != nil {
		return err
	}

	return nil
}
