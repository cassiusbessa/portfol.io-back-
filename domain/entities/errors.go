package entities

import "fmt"

func EmptyFieldError(field string) error {
	return fmt.Errorf("\"%s\" está vazio", field)
}

func InvalidFieldLengthError(field string, min int, max int) error {
	return fmt.Errorf("\"%s\" tem tamanho inválido, o campo deve ter entre min: %d e max: %d caracteres", field, min, max)
}

func NegativeIntegerError(field string) error {
	return fmt.Errorf("\"%s\" não pode ser negativo", field)
}

func InvalidEmailError() error {
	return fmt.Errorf("email inválido")
}

func InvalidPatternError(field, rule string) error {
	return fmt.Errorf("\"%s\" deve ter pelo menos um  %s caracter", field, rule)
}

func InvalidFieldRangeError(field string, min int, max int) error {
	return fmt.Errorf("\"%s\" o campo deve ter entre min: %d, max: %d", field, min, max)
}
