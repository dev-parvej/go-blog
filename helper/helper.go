package helper

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/mitchellh/mapstructure"
)

func ParseForm(r *http.Request, output interface{}) interface{} {
	r.ParseForm()

	form := map[string]string{}
	for index, value := range r.Form {
		form[index] = value[0]
	}

	return mapstructure.Decode(form, output)
}

func ValidateForm(form interface{}) error {
	err := validator.New().Struct(form)
	if err == nil {
		return nil
	}

	errors := err.(validator.ValidationErrors)

	if errors != nil {
		return errors
	}

	return nil

}
