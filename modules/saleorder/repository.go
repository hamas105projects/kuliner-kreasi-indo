package saleorder

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SaleOrderRepository struct {
	DB *gorm.DB
}

func (r *SaleOrderRepository) Create(order *SaleOrder) error {
	tx := r.DB.Begin()
	if err := tx.Create(order).Error; err != nil {
		tx.Rollback()
		return err
	}

	for _, item := range order.Items {
		if err := tx.Create(&item).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func (r *SaleOrderRepository) GetByID(id string) (SaleOrder, error) {
	var order SaleOrder
	err := r.DB.Preload("Items").First(&order, "id = ?", id).Error
	return order, err
}

func (r *SaleOrderRepository) GetAll(offset, limit int) ([]SaleOrder, error) {
	var orders []SaleOrder
	err := r.DB.Preload("Items").Offset(offset).Limit(limit).Find(&orders).Error
	return orders, err
}

func (r *SaleOrderRepository) Update(order *SaleOrder) error {
	tx := r.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// Update SaleOrder utama
	if err := tx.Model(&SaleOrder{}).
		Where("id = ?", order.ID).
		Updates(map[string]interface{}{
			"total_amount": order.TotalAmount,
			"updated_at":   time.Now(),
		}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Hapus semua item lama
	if err := tx.Where("sale_order_id = ?", order.ID).
		Delete(&SaleOrderItem{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Insert item baru satu per satu (antrian)
	for _, item := range order.Items {
		item.ID = uuid.New()
		item.SaleOrderID = order.ID
		if err := tx.Create(&item).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func (r *SaleOrderRepository) Delete(id string) error {
	return r.DB.Delete(&SaleOrder{}, "id = ?", id).Error
}
