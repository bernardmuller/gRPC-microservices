package main

import (
	"github.com/bernardmuller/gRPC-microservices/services/kitchen/infrastructure"
)

func main() {
	httpServer := infrastructure.NewHttpServer(":1000")
	httpServer.Run()
}
