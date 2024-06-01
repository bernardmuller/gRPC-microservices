package infrastructure

import (
	"log"
	"net/http"

	handler "github.com/bernardmuller/gRPC-microservices/services/orders/handler/orders"
	"github.com/bernardmuller/gRPC-microservices/services/orders/service"
)

type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{addr: addr}
}

func (s *httpServer) Start() error {
	router := http.NewServeMux()

	orderService := service.NewOrdersService()
	orderHandler := handler.NewHttpOrdersHandler(orderService)
	orderHandler.RegisterRouter(router)

	log.Println("Starting server on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
