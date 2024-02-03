package entities

import "regexp"

func emptyStringValidator(value, name string) error {
	if value == "" {
		return EmptyFieldError(name)
	}
	return nil
}

func invalidFieldLenghValidator(field string, name string, min int, max int) error {
	if len(field) < min || len(field) > max {
		return InvalidFieldLengthError(name, min, max)
	}
	return nil
}

func invalidRangeIntValidator(field int, name string, min int, max int) error {
	if field < min || field > max {
		return InvalidFieldRangeError(name, min, max)
	}
	return nil
}

func emailPatternValidator(email string) error {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)

	if !re.MatchString(email) {
		return InvalidEmailError()
	}
	return nil
}

func passwordPatternValidator(password string) error {
	lowercaseRegex := `.*[a-z].*`
	uppercaseRegex := `.*[A-Z].*`
	digitRegex := `.*\d.*`

	if !regexp.MustCompile(lowercaseRegex).MatchString(password) {
		return InvalidPatternError("Senha", "pelo menos uma letra minúscula")
	}
	if !regexp.MustCompile(uppercaseRegex).MatchString(password) {
		return InvalidPatternError("Senha", "pelo menos uma letra maiúscula")
	}
	if !regexp.MustCompile(digitRegex).MatchString(password) {
		return InvalidPatternError("Senha", "pelo menos um dígito numérico")
	}
	return nil
}
