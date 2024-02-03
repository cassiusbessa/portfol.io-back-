package usecases

import "fmt"

func entityAlreadyExists(entity string, field string, value string) error {
	return fmt.Errorf("%s já possui um registro do campo %s com o valor: %s", entity, field, value)
}

func entityNotFound(entity string, field string, value string) error {
	return fmt.Errorf("%s não encontrado(a) com o campo %s com o valor: %v", entity, field, value)
}

func invalidLength(field string, min int, max int) error {
	return fmt.Errorf("%s deve ter entre %d e %d registros", field, min, max)
}
