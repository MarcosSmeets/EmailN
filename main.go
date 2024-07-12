package main

import (
	"emailn/internal/domain/campaing"

	"github.com/go-playground/validator/v10"
)

func main() {
	campaing := campaing.Campaing{}

	validate := validator.New()
	err := validate.Struct(campaing)
	if err == nil {
		println("no error")
	} else {
		validationErrors := err.(validator.ValidationErrors)
		for _, v := range validationErrors {
			switch v.Tag() {
			case "required":
				println(v.StructField() + " is required")
			case "min":
				println(v.StructField() + " is required with min " + v.Param())
			case "max":
				println(v.StructField() + " is required with max " + v.Param())
			case "email":
				println(v.StructField() + " is invalid " + v.Param())
			}
		}
	}
}
