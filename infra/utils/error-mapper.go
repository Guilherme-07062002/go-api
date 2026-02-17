package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func TranslateError(err error) map[string]string {
	errors := make(map[string]string)

	validationErrors, ok := err.(validator.ValidationErrors)
	if ok == false {
		return map[string]string{"error": "Erro de validação desconhecido"}
	}

	for _, f := range validationErrors {
		switch f.Tag() {
		case "required":
			errors[f.Field()] = "Este campo é obrigatório"
		case "min":
			errors[f.Field()] = fmt.Sprintf("O valor mínimo é %s", f.Param())
		case "max":
			errors[f.Field()] = fmt.Sprintf("O valor máximo é %s", f.Param())
		case "gt":
			errors[f.Field()] = fmt.Sprintf("O valor deve ser maior que %s", f.Param())
		case "email":
			errors[f.Field()] = "E-mail inválido"
		default:
			errors[f.Field()] = fmt.Sprintf("Falha na validação: %s", f.Tag())
		}
	}

	return errors
}
