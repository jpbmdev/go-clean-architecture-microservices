# go-clean-architecture-microservices

# generate swagger docs

`swag init -g infrastructure/docs/customer.docs.go controllers/customer.controller.go`

# generate unit testing coverage

`go test -coverprofile coverage.out ./...`

Do not run this on a power shell terminal
`go tool cover -html=coverage.out -o coverage.html`

# debugger config

https://dev.to/andreidascalu/setup-go-with-vscode-in-docker-for-debugging-24ch
