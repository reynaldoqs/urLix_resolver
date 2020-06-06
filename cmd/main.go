package main

import "github.com/reynaldoqs/urLix_resolver/internal/infrastructure/server"

func main() {
	server.RegisterRouter(":8080")
}
