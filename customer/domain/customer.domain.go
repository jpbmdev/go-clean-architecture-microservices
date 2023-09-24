package domain

import "github.com/go-playground/validator/v10"

type CustomerEntity struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Age       int    `json:"age" validate:"required,gt=17"`
	Score     int    `json:"score" validate:"required"`
	Id        uint   `json:"id"`
}

type CreateCustomerDto struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Age       int    `json:"age" validate:"required,gt=17"`
	Score     int    `json:"score" validate:"required"`
}

func ValidateCreateCustomerDto(createCustomerDto *CreateCustomerDto) string {
	validate := validator.New()

	err := validate.Struct(createCustomerDto)

	if err == nil {
		return ""
	}

	// If there are validation errors, concatenate them into a string
	errorMsg := ""
	for _, e := range err.(validator.ValidationErrors) {
		errorMsg += e.Field() + ": " + e.Tag() + "; "
	}
	return errorMsg
}
