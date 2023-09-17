package server

import (
	"fmt"

	"go-clean-architecture-mircroservices_customer/core/database"
	"go-clean-architecture-mircroservices_customer/infrastructure/docs"
	"go-clean-architecture-mircroservices_customer/infrastructure/models"
	"go-clean-architecture-mircroservices_customer/infrastructure/routes"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	port string
	app  *fiber.App
}

func (s *Server) Initialize(port string) {
	database.ConnectToDB()

	database.DBconnection.AutoMigrate(&models.CustomerModel{})

	s.port = port
	s.app = fiber.New()

	routes.AddCustomerRoutesToApp(s.app)

	docs.GenerateDocsRoute(s.app)
}

func (s *Server) Listen() {
	s.app.Listen(s.port)

	fmt.Print("Server listening on port: " + s.port)

	db, _ := database.DBconnection.DB()
	defer db.Close()
}
