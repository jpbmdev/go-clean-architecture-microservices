package adapters

import (
	interfaces "go-clean-architecture-mircroservices_customer/core/interfaces"

	"github.com/gofiber/fiber/v2"
)

func FiberAdapter(httpHandler interfaces.HttpHandler) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		response := httpHandler(interfaces.HttpRequest{
			Params: ctx.AllParams(),
			Body:   ctx.Body(),
		})

		return ctx.Status(response.StausCode).JSON(response.Body)
	}
}
