package usecases

import "fmt"

func entityAlreadyExists(entity string, fields string, values map[string]string) error {
	return fmt.Errorf("%s already exists with %s: %v", entity, fields, values)
}

func entityNotFound(entity string, fields string, values map[string]string) error {
	return fmt.Errorf("%s not found with %s: %v", entity, fields, values)
}
