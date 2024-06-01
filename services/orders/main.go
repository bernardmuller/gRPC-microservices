package main

import (
	"github.com/bernardmuller/gRPC-microservices/services/orders/infrastructure"
)

func main() {
	httpServer := infrastructure.NewHttpServer(":8000")
	go httpServer.Start()

	grpcServer := infrastructure.NewGrpcServer(":9000")
	grpcServer.Start()
}
