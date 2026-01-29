package main

import (
	"pizza-tracker/internal/models"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func RegisterCustomValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("valid_pizza_type", createSliceValidator(models.PizzaTypes))
		v.RegisterValidation("valid_pizza_size", createSliceValidator(models.PizzaSizes))
		v.RegisterValidation("valid_order_status", createSliceValidator(models.OrderStatuses))
	}
}

func createSliceValidator(allowValues []string) validator.Func {
	return func(fl validator.FieldLevel) bool {
		for _, value := range allowValues {
			if value == fl.Field().String() {
				return true
			}
		}
		return false
	}
}
