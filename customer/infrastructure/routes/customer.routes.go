package routes

import (
	"go-clean-architecture-mircroservices_customer/adapters"
	"go-clean-architecture-mircroservices_customer/controllers"
	"go-clean-architecture-mircroservices_customer/services"

	"github.com/gofiber/fiber/v2"
)

func AddCustomerRoutesToApp(app *fiber.App) {
	service := services.CreateCustomerService()
	controller := controllers.CustomerController{CustomerService: service}

	app.Get("/customer", adapters.FiberAdapter(controller.GetCustomer))

	app.Post("/customer", adapters.FiberAdapter(controller.CreateCustomer))
}
