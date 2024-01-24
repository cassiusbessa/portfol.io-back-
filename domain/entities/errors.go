package entities

import "fmt"

func OnlyOneCpfOrCnpjError() error {
	return fmt.Errorf("only one should be filed, CPF or CNPJ, not both")
}

func EmptyCpfAndCnpjError() error {
	return fmt.Errorf("at least one of the fields, CPF or CNPJ, must be filled in")
}

func EmptyFieldError(field string) error {
	return fmt.Errorf("\"%s\" is empty", field)
}

func InvalidFieldPatternError(field string) error {
	return fmt.Errorf("\"%s\" has invalid pattern", field)
}

func InvalidFieldLengthError(field string, min int, max int) error {
	return fmt.Errorf("\"%s\" has invalid length, this field must be between min: %d, max: %d characteres", field, min, max)
}

func NegativeIntegerError(field string) error {
	return fmt.Errorf("\"%s\" cannot be negative", field)
}

func InvalidEmailError() error {
	return fmt.Errorf("invalid email")
}

func InvalidPatternError(field, rule string) error {
	return fmt.Errorf("\"%s\" must have at least one %s character", field, rule)
}

func InvalidFieldRangeError(field string, min int, max int) error {
	return fmt.Errorf("\"%s\" has invalid range, this field must be between min: %d, max: %d", field, min, max)
}
