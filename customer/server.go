package main

import (
	"fmt"

	infrastructure "go-clean-architecture-mircroservices_customer/infrastructure"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	port string
	app  *fiber.App
}

func (s *Server) Initialize(port string) {
	s.port = port
	s.app = fiber.New()

	infrastructure.AddCustomerRoutesToApp(s.app)
}

func (s *Server) Listen() {
	s.app.Listen(s.port)
	fmt.Print("Server listening on port: " + s.port)
}
