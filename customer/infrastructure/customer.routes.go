package infrastructure

import (
	"go-clean-architecture-mircroservices_customer/adapters"
	"go-clean-architecture-mircroservices_customer/controller"

	"github.com/gofiber/fiber/v2"
)

func AddCustomerRoutesToApp(app *fiber.App) {
	controller := controller.CustomerController{}

	app.Get("/customer", adapters.FiberAdapter(controller.GetCustomer))

	app.Post("/customer", adapters.FiberAdapter(controller.CreateCustomer))
}
