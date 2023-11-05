package mock

import (
	"errors"
	"go-clean-architecture-mircroservices_customer/domain"
)

type MockCustomerRepository struct{}

func (cr *MockCustomerRepository) GetCustomers() ([]domain.CustomerEntity, error) {
	customers := []domain.CustomerEntity{{Id: 1, FirstName: "1", LastName: "2", Score: 10, Age: 18}}
	return customers, nil
}
func (cr *MockCustomerRepository) CreateCustomer(customerDto domain.CreateCustomerDto) (domain.CustomerEntity, error) {
	return domain.CustomerEntity{
		Id: 1, FirstName: "1", LastName: "2", Score: 10, Age: 18,
	}, nil
}

type MockCustomerErrorRepository struct{}

func (cr *MockCustomerErrorRepository) GetCustomers() ([]domain.CustomerEntity, error) {
	customers := []domain.CustomerEntity{}
	return customers, errors.New("db failed")
}
func (cr *MockCustomerErrorRepository) CreateCustomer(customerDto domain.CreateCustomerDto) (domain.CustomerEntity, error) {
	return domain.CustomerEntity{}, errors.New("db failed")
}
