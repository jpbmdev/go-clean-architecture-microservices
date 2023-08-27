package main

import (
	"go-clean-architecture-mircroservices_customer/core/server"
)

func main() {
	server := server.Server{}

	server.Initialize(":3000")

	server.Listen()
}
