package services

import (
	"fmt"
	"regexp"
)

const (
	emailRegex = `[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`
	cpfRegex   = `^\d{3}\.?\d{3}\.?\d{3}-?\d{2}$`
)

// ValidateEmail verifica se o email fornecido está em um formato válido
func ValidateEmail(email string) (bool, error) {
	match, err := regexp.MatchString(emailRegex, email)
	if err != nil {
		return false, fmt.Errorf("erro to validate email: %s", err.Error())
	}
	return match, nil
}

// ValidateCPF verifica se o CPF fornecido está em um formato válido
func ValidateCPF(CPF string) (bool, error) {
	match, err := regexp.MatchString(cpfRegex, CPF)
	if err != nil {
		return false, fmt.Errorf("error to validate CPF: %s", err.Error())
	}
	return match, nil
}
