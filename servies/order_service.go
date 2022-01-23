package servies

import (
	"eshop/datamodels"
	"eshop/repositories"
)

type IOrderService interface {
	GetOrderById(int64) (*datamodels.Order, error)
	GetAllOrder() ([]*datamodels.Order, error)
	DeleteOrderById(int64) bool
	InsertOrder(*datamodels.Order) (int64, error)
	UpdateOrder(*datamodels.Order) error
	GetAllOrderInfo() (map[int]map[string]string, error)
}

type OrderService struct {
	orderRepository repositories.IOrder
}

func NewOrderService(orderRepository repositories.IOrder) IOrderService {
	return &OrderService{orderRepository}
}

func (s *OrderService) GetOrderById(key int64) (*datamodels.Order, error) {
	return s.orderRepository.SelectByKey(key)
}

func (s *OrderService) GetAllOrder() ([]*datamodels.Order, error) {
	return s.orderRepository.SelectAll()
}

func (s *OrderService) DeleteOrderById(key int64) bool {
	return s.orderRepository.Delete(key)
}

func (s *OrderService) InsertOrder(order *datamodels.Order) (int64, error) {
	return s.orderRepository.Insert(order)
}

func (s *OrderService) UpdateOrder(order *datamodels.Order) error {
	return s.orderRepository.Update(order)
}

func (s *OrderService) GetAllOrderInfo() (map[int]map[string]string, error) {
	return s.orderRepository.SelectAllWithInfo()
}
