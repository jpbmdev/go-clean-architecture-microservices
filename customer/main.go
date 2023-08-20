package main

func main() {
	server := Server{}

	server.Initialize(":3000")

	server.Listen()
}
