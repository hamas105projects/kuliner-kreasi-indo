package saleorder

import (
	"errors"

	"github.com/google/uuid"
)

type SaleOrderService struct {
	Repo *SaleOrderRepository
}

func (s *SaleOrderService) CreateSaleOrder(cashierID string, items []SaleOrderItem) (SaleOrder, error) {
	order := SaleOrder{
		ID:        uuid.New(),
		CashierID: uuid.MustParse(cashierID),
		Items:     items,
	}

	total := 0.0
	for _, item := range items {
		total += float64(item.Qty) * item.Price
	}
	order.TotalAmount = total

	err := s.Repo.Create(&order)
	return order, err
}

func (s *SaleOrderService) GetAllSaleOrders(offset, limit int) ([]SaleOrder, error) {
	return s.Repo.GetAll(offset, limit)
}

func (s *SaleOrderService) GetSaleOrderByID(id string) (SaleOrder, error) {
	order, err := s.Repo.GetByID(id)
	if err != nil {
		return order, errors.New("Sale order not found")
	}
	return order, nil
}

func (s *SaleOrderService) UpdateSaleOrder(id string, items []SaleOrderItem) (SaleOrder, error) {
	order, err := s.Repo.GetByID(id)
	if err != nil {
		return order, errors.New("Sale order not found")
	}

	order.Items = items
	total := 0.0
	for _, item := range items {
		total += float64(item.Qty) * item.Price
	}
	order.TotalAmount = total

	err = s.Repo.Update(&order)
	return order, err
}

func (s *SaleOrderService) DeleteSaleOrder(id string) error {
	return s.Repo.Delete(id)
}
