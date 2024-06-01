package infrastructure

import (
	handler "github.com/bernardmuller/gRPC-microservices/services/orders/handler/orders"
	"github.com/bernardmuller/gRPC-microservices/services/orders/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	addr string
}

func NewGrpcServer(addr string) *Server {
	return &Server{addr: addr}
}

func (s *Server) Start() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return err
	}

	grpcServer := grpc.NewServer()

	orderService := service.NewOrdersService()
	handler.NewGRPCOrdersService(grpcServer, orderService)

	log.Println("Server is running on port", s.addr)

	return grpcServer.Serve(lis)
}
