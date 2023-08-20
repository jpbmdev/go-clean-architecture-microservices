package controller

import (
	"encoding/json"
	interfaces "go-clean-architecture-mircroservices_customer/core/interfaces"
	domain "go-clean-architecture-mircroservices_customer/domain"
)

type CustomerController struct{}

func (cc *CustomerController) GetCustomer(httpRequest interfaces.HttpRequest) interfaces.HttpResponse {
	customer := domain.CustomerEntity{
		Name: "PEPE",
	}

	return interfaces.HttpResponse{StausCode: 200, Body: customer}
}

func (cc *CustomerController) CreateCustomer(httpRequest interfaces.HttpRequest) interfaces.HttpResponse {

	var customer domain.CustomerEntity

	if err := json.Unmarshal(httpRequest.Body, &customer); err != nil {
		return interfaces.HttpResponse{StausCode: 400, Body: interfaces.Error{Message: "Invalid body"}}
	}

	return interfaces.HttpResponse{StausCode: 200, Body: customer}
}
