package controllers

import (
	"encoding/json"
	"log"
	"testing"

	"go-clean-architecture-mircroservices_customer/core/interfaces"
	"go-clean-architecture-mircroservices_customer/domain"
	"go-clean-architecture-mircroservices_customer/mock"
	"go-clean-architecture-mircroservices_customer/services"
)

var customerToCreate = domain.CreateCustomerDto{FirstName: "1", LastName: "2", Email: "test@test.com", Score: 10, Age: 18}

// GetCustomer ------------------------------

func TestGetCustomersOk(t *testing.T) {
	var repository = &mock.MockCustomerRepository{}
	var service = services.CustomerService{Repository: repository}
	var controller = CustomerController{Service: service}

	httpRequest := interfaces.HttpRequest{}
	response := controller.GetCustomers(httpRequest)

	data, ok := response.Body.([]domain.CustomerEntity)

	if !ok {
		t.Errorf("The response is not an array of CustomerEntity")
	}

	currentLenght := len(data)
	expectedLenght := 1
	if currentLenght != expectedLenght {
		t.Errorf("Expected %d, but got %d", expectedLenght, currentLenght)
	}
}

func TestGetCustomersError(t *testing.T) {
	var repository = &mock.MockCustomerErrorRepository{}
	var service = services.CustomerService{Repository: repository}
	var controller = CustomerController{Service: service}

	httpRequest := interfaces.HttpRequest{}
	response := controller.GetCustomers(httpRequest)

	data, ok := response.Body.(interfaces.Error)

	if !ok {
		t.Errorf("The response is not an error")
	}

	expectedError := "Failed to get customers"
	if data.Message != expectedError {
		t.Errorf("Expected %s error but got %s", expectedError, data.Message)
	}
}

// CreateCustomer ------------------------------

func TestCreateCustomerOk(t *testing.T) {
	var repository = &mock.MockCustomerRepository{}
	var service = services.CustomerService{Repository: repository}
	var controller = CustomerController{Service: service}

	byteSlice, err := json.Marshal(customerToCreate)
	if err != nil {
		log.Fatal("Marshal error:", err)
	}

	httpRequest := interfaces.HttpRequest{Body: byteSlice}

	response := controller.CreateCustomer(httpRequest)

	expectedStatus := 201

	if response.StausCode != expectedStatus {
		t.Errorf("Expected %d status, but got %d", expectedStatus, response.StausCode)
	}
}

func TestCreateCustomerBodyError(t *testing.T) {
	var repository = &mock.MockCustomerRepository{}
	var service = services.CustomerService{Repository: repository}
	var controller = CustomerController{Service: service}

	httpRequest := interfaces.HttpRequest{}

	response := controller.CreateCustomer(httpRequest)

	data, ok := response.Body.(interfaces.Error)
	if !ok {
		t.Errorf("The response is not an error")
	}

	expectedStatus := 400
	if response.StausCode != expectedStatus {
		t.Errorf("Expected %d status, but got %d", expectedStatus, response.StausCode)
	}

	expectedMessage := "Invalid body"
	if data.Message != expectedMessage {
		t.Errorf("Expected %s message, but got %s", expectedMessage, data.Message)
	}
}

func TestCreateCustomerDtoError(t *testing.T) {
	var repository = &mock.MockCustomerRepository{}
	var service = services.CustomerService{Repository: repository}
	var controller = CustomerController{Service: service}

	badCustomerToCreate := domain.CreateCustomerDto{FirstName: "1", LastName: "2", Score: 10, Age: 18}

	byteSlice, err := json.Marshal(badCustomerToCreate)
	if err != nil {
		log.Fatal("Marshal error:", err)
	}

	httpRequest := interfaces.HttpRequest{Body: byteSlice}

	response := controller.CreateCustomer(httpRequest)

	data, ok := response.Body.(interfaces.Error)
	if !ok {
		t.Errorf("The response is not an error")
	}

	expectedStatus := 400
	if response.StausCode != expectedStatus {
		t.Errorf("Expected %d status, but got %d", expectedStatus, response.StausCode)
	}

	expectedMessage := "Email: required; "
	if data.Message != expectedMessage {
		t.Errorf("Expected %s message, but got %s", expectedMessage, data.Message)
	}
}

func TestCreateCustomerDBError(t *testing.T) {
	var repository = &mock.MockCustomerErrorRepository{}
	var service = services.CustomerService{Repository: repository}
	var controller = CustomerController{Service: service}

	byteSlice, err := json.Marshal(customerToCreate)
	if err != nil {
		log.Fatal("Marshal error:", err)
	}

	httpRequest := interfaces.HttpRequest{Body: byteSlice}

	response := controller.CreateCustomer(httpRequest)

	expectedStatus := 500
	if response.StausCode != expectedStatus {
		t.Errorf("Expected %d status, but got %d", expectedStatus, response.StausCode)
	}
}
