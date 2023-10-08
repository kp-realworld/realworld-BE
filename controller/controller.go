package controller

import "github.com/go-playground/validator/v10"

func init() {
	Validate = validator.New()
}
