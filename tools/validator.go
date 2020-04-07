package tools

import (
	"fmt"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"strings"
)

// Validator struct
type Validator struct {
	Validator  *validator.Validate
	Translator ut.Translator
	Error      validator.ValidationErrors
}

// Validate function to validate
func (v *Validator) Validate(i interface{}) error {
	var (
		errorList []string
		error     error
	)

	err := v.Validator.Struct(i)
	if err != nil {
		v.Error = err.(validator.ValidationErrors)
		errorList = v.Errors()
	}

	if len(errorList) > 0 {
		error = fmt.Errorf(strings.Join(errorList, "\n"))
	}

	return error
}

func (v *Validator) Errors() []string {
	var errString = make([]string, 0)

	for _, err := range v.Error {
		errString = append(errString, fmt.Sprintf("%s|%s", strings.ToLower(err.Field()), err.Translate(v.Translator)))
	}

	return errString
}
