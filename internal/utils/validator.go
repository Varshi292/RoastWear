package utils

import (
	"github.com/go-playground/validator/v10"
	"strconv"
)

var Validate *validator.Validate

func InitializeValidator() error {
	Validate = validator.New()
	if err := Validate.RegisterValidation("port", validatePort); err != nil {
		return err
	}
	return nil
}

func validatePort(fl validator.FieldLevel) bool {
	port := fl.Field().String()
	portInt, err := strconv.Atoi(port)
	return err == nil && portInt > 0 && portInt <= 65535
}
