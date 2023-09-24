package repository

import (
	"fmt"
	"go-clean-architecture-mircroservices_customer/core/database"
	"go-clean-architecture-mircroservices_customer/domain"
	"go-clean-architecture-mircroservices_customer/infrastructure/models"
)

type customerRepository struct{}

func CreateCustomerRepository() domain.CustomerRepository {
	return &customerRepository{}
}

func (cr *customerRepository) GetCustomers() ([]domain.CustomerEntity, error) {
	var models []models.CustomerModel
	var customers []domain.CustomerEntity

	response := database.DBconnection.Find(&models)
	if response.Error != nil {
		return customers, response.Error
	}

	for _, value := range models {
		customers = append(customers, domain.CustomerEntity{
			Id:        value.Id,
			FirstName: value.FirstName,
			LastName:  value.LastName,
			Email:     value.Email,
			Age:       value.Age,
			Score:     value.Score,
		})
	}

	return customers, nil
}

func (cr *customerRepository) CreateCustomer(customerDto domain.CreateCustomerDto) (domain.CustomerEntity, error) {
	model := models.CustomerModel{
		FirstName: customerDto.FirstName,
		LastName:  customerDto.LastName,
		Email:     customerDto.Email,
		Age:       customerDto.Age,
		Score:     customerDto.Score,
	}
	createdUser := domain.CustomerEntity{}

	response := database.DBconnection.Create(&model)
	if response.Error != nil {
		return createdUser, response.Error
	}

	fmt.Print(model)

	return domain.CustomerEntity{
		Id:        model.Id,
		FirstName: model.FirstName,
		LastName:  model.LastName,
		Email:     model.Email,
		Age:       model.Age,
		Score:     model.Score,
	}, nil
}
