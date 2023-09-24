package routes

import (
	"go-clean-architecture-mircroservices_customer/adapters"
	"go-clean-architecture-mircroservices_customer/controllers"
	"go-clean-architecture-mircroservices_customer/infrastructure/repository"
	"go-clean-architecture-mircroservices_customer/services"

	"github.com/gofiber/fiber/v2"
)

func AddCustomerRoutesToApp(app *fiber.App) {
	repository := repository.CreateCustomerRepository()
	service := services.CustomerService{Repository: repository}
	controller := controllers.CustomerController{Service: service}

	app.Get("/customer", adapters.FiberAdapter(controller.GetCustomers))

	app.Post("/customer", adapters.FiberAdapter(controller.CreateCustomer))
}
