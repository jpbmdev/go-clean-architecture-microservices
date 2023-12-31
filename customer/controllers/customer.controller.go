package controllers

import (
	"encoding/json"
	"go-clean-architecture-mircroservices_customer/core/interfaces"
	"go-clean-architecture-mircroservices_customer/domain"
	"go-clean-architecture-mircroservices_customer/services"
)

type CustomerController struct {
	Service services.CustomerService
}

// @Summary Get Customer
// @Description Get the list of customers
// @Tags Customers
// @ID get-customer-by-id
// @Accept json
// @Produce json
// @Success 200 {object} domain.CustomerEntity
// @Failure 500 {object} interfaces.Error
// @Router /customer [get]
func (cc *CustomerController) GetCustomers(httpRequest interfaces.HttpRequest) interfaces.HttpResponse {
	customers, err := cc.Service.GetCustomers()
	if err != nil {
		return interfaces.HttpResponse{StausCode: 500, Body: interfaces.Error{Message: "Failed to get customers"}}
	}

	return interfaces.HttpResponse{StausCode: 200, Body: customers}
}

// @Summary Create a customer
// @Description Create a new customer with the provided name.
// @Tags Customers
// @ID create-customer
// @Accept json
// @Produce json
// @Param request body domain.CreateCustomerDto true "Customer details to create"
// @Success 201 {object} domain.CustomerEntity
// @Failure 400 {object} interfaces.Error
// @Failure 500 {object} interfaces.Error
// @Router /customer [post]
func (cc *CustomerController) CreateCustomer(httpRequest interfaces.HttpRequest) interfaces.HttpResponse {

	var customerDto domain.CreateCustomerDto

	if err := json.Unmarshal(httpRequest.Body, &customerDto); err != nil {
		return interfaces.HttpResponse{StausCode: 400, Body: interfaces.Error{Message: "Invalid body"}}
	}

	if errMsg := domain.ValidateCreateCustomerDto(&customerDto); errMsg != "" {
		return interfaces.HttpResponse{StausCode: 400, Body: interfaces.Error{Message: errMsg}}
	}

	customer, err := cc.Service.CreateCustomer(customerDto)
	if err != nil {
		return interfaces.HttpResponse{StausCode: 500, Body: interfaces.Error{Message: "Failed to create Customer"}}
	}

	return interfaces.HttpResponse{StausCode: 201, Body: customer}
}
