package services

import (
	"fmt"
	"go-clean-architecture-mircroservices_customer/core/database"
	"go-clean-architecture-mircroservices_customer/domain"
	"go-clean-architecture-mircroservices_customer/infrastructure/models"
)

type CustomerService interface {
	GetCustomers() ([]domain.CustomerEntity, error)
	CreateCustomer(domain.CreateCustomerDto) (domain.CustomerEntity, error)
}

type customerService struct{}

func CreateCustomerService() CustomerService {
	return &customerService{}
}

func (s *customerService) GetCustomers() ([]domain.CustomerEntity, error) {
	var models []models.CustomerModel
	var customers []domain.CustomerEntity

	response := database.DBconnection.Find(&models)
	if response.Error != nil {
		return customers, response.Error
	}

	for _, value := range models {
		customers = append(customers, domain.CustomerEntity{
			Id:   value.Id,
			Name: value.Name,
		})
	}

	return customers, nil
}

func (s *customerService) CreateCustomer(customerDto domain.CreateCustomerDto) (domain.CustomerEntity, error) {
	model := models.CustomerModel{
		Name: customerDto.Name,
	}
	createdUser := domain.CustomerEntity{}

	response := database.DBconnection.Create(&model)
	if response.Error != nil {
		return createdUser, response.Error
	}

	fmt.Print(model)

	return domain.CustomerEntity{
		Id:   model.Id,
		Name: model.Name,
	}, nil
}
