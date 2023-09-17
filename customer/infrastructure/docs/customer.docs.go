package docs

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	_ "go-clean-architecture-mircroservices_customer/docs"
)

// @title Swagger Example API
// @version 1.0
// @description This is test with clean architecture
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
func GenerateDocsRoute(app *fiber.App) {
	app.Get("/docs/*", swagger.HandlerDefault)
}
