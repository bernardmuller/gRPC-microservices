package handler

import (
	"context"
	"github.com/bernardmuller/gRPC-microservices/services/common/genproto/orders"
	"github.com/bernardmuller/gRPC-microservices/services/orders/types"
	"google.golang.org/grpc"
)

type OrdersHandler struct {
	ordersService types.OrderService
	orders.UnimplementedOrderServiceServer
}

func NewGRPCOrdersService(grpc *grpc.Server, ordersService types.OrderService) {
	grpcHandler := &OrdersHandler{
		ordersService: ordersService,
	}

	orders.RegisterOrderServiceServer(grpc, grpcHandler)
}

func (h *OrdersHandler) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {
	order := &orders.Order{
		OrderID:    42,
		ProductID:  1,
		Quantity:   10,
		CustomerID: 12,
	}

	err := h.ordersService.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	res := &orders.CreateOrderResponse{
		Status: "success",
	}

	return res, nil
}

func (h *OrdersHandler) GetOrders(ctx context.Context, req *orders.GetOrdersRequest) (*orders.GetOrderResponse, error) {
	o := h.ordersService.GetOrders(ctx)
	res := &orders.GetOrderResponse{
		Orders: o,
	}

	return res, nil
}
