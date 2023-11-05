package services

import (
	"go-clean-architecture-mircroservices_customer/domain"
	"go-clean-architecture-mircroservices_customer/mock"
	"testing"
)

var customerToCreate = domain.CreateCustomerDto{FirstName: "1", LastName: "2", Email: "test@test.com", Score: 10, Age: 18}

func TestGetCustomersOk(t *testing.T) {
	var repository = &mock.MockCustomerRepository{}
	var service = CustomerService{Repository: repository}

	response, err := service.GetCustomers()

	if err != nil {
		t.Errorf("Failed to Get Customers")
	}

	currentLenght := len(response)
	expectedLenght := 1
	if currentLenght != expectedLenght {
		t.Errorf("Expected %d, but got %d", expectedLenght, currentLenght)
	}
}

func TestCreateCustomerOk(t *testing.T) {
	var repository = &mock.MockCustomerRepository{}
	var service = CustomerService{Repository: repository}

	_, err := service.CreateCustomer(customerToCreate)

	if err != nil {
		t.Errorf("Failed to Get Customers")
	}

}
