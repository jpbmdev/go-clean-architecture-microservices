package services

import "go-clean-architecture-mircroservices_customer/domain"

type CustomerService struct {
	Repository domain.CustomerRepository
}

func (cs *CustomerService) GetCustomers() ([]domain.CustomerEntity, error) {
	return cs.Repository.GetCustomers()
}

func (cs *CustomerService) CreateCustomer(customerDto domain.CreateCustomerDto) (domain.CustomerEntity, error) {
	return cs.Repository.CreateCustomer(customerDto)
}
