package saleorder

import (
	"errors"

	"github.com/google/uuid"
)

type SaleOrderService struct {
	Repo *SaleOrderRepository
}

func (s *SaleOrderService) CreateSaleOrder(cashierID string, items []SaleOrderItem) (SaleOrder, error) {
	// Mulai transaction
	tx := s.Repo.DB.Begin()
	if tx.Error != nil {
		return SaleOrder{}, tx.Error
	}

	order := SaleOrder{
		ID:        uuid.New(),
		CashierID: uuid.MustParse(cashierID),
	}

	total := 0.0

	// Simpan SaleOrder dulu
	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		return order, err
	}

	// Insert item satu per satu (antrian)
	for _, item := range items {
		item.ID = uuid.New()
		item.SaleOrderID = order.ID

		// Ambil harga produk dari tabel products (snapshot)
		var product Product
		if err := tx.First(&product, "id = ?", item.ProductID).Error; err != nil {
			tx.Rollback()
			return order, err
		}
		item.PriceSnapshot = product.Price

		// Hitung total
		total += float64(item.Qty) * item.PriceSnapshot

		// Insert item
		if err := tx.Create(&item).Error; err != nil {
			tx.Rollback()
			return order, err
		}
	}

	// Update total order
	order.TotalAmount = total
	if err := tx.Save(&order).Error; err != nil {
		tx.Rollback()
		return order, err
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return order, err
	}

	return order, nil
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
		total += float64(item.Qty) * item.PriceSnapshot
	}
	order.TotalAmount = total

	err = s.Repo.Update(&order)
	return order, err
}

func (s *SaleOrderService) DeleteSaleOrder(id string) error {
	return s.Repo.Delete(id)
}
