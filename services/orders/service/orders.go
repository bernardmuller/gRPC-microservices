package service

import (
	"context"
	"github.com/bernardmuller/gRPC-microservices/services/common/genproto/orders"
)

var ordersDb = make([]*orders.Order, 0)

type OrdersService struct {
}

func NewOrdersService() *OrdersService {
	return &OrdersService{}
}

func (s *OrdersService) CreateOrder(ctx context.Context, order *orders.Order) error {
	ordersDb = append(ordersDb, order)
	return nil
}

func (s *OrdersService) GetOrders(ctx context.Context) []*orders.Order {
	return ordersDb
}
