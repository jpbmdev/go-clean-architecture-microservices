package controllers

import (
	"encoding/json"
	"go-clean-architecture-mircroservices_customer/core/interfaces"
	"go-clean-architecture-mircroservices_customer/domain"
	"go-clean-architecture-mircroservices_customer/services"
)

type CustomerController struct {
	CustomerService services.CustomerService
}

func (cc *CustomerController) GetCustomers(httpRequest interfaces.HttpRequest) interfaces.HttpResponse {
	customers, err := cc.CustomerService.GetCustomers()
	if err != nil {
		return interfaces.HttpResponse{StausCode: 500, Body: interfaces.Error{Message: "Failed to get customers"}}
	}

	return interfaces.HttpResponse{StausCode: 200, Body: customers}
}

func (cc *CustomerController) CreateCustomer(httpRequest interfaces.HttpRequest) interfaces.HttpResponse {

	var customerDto domain.CreateCustomerDto

	if err := json.Unmarshal(httpRequest.Body, &customerDto); err != nil {
		return interfaces.HttpResponse{StausCode: 400, Body: interfaces.Error{Message: "Invalid body"}}
	}

	customer, err := cc.CustomerService.CreateCustomer(customerDto)
	if err != nil {
		return interfaces.HttpResponse{StausCode: 500, Body: interfaces.Error{Message: "Failed to create Customer"}}
	}

	return interfaces.HttpResponse{StausCode: 200, Body: customer}
}
