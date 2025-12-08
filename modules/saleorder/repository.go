package saleorder

import "gorm.io/gorm"

type SaleOrderRepository struct {
	DB *gorm.DB
}

func (r *SaleOrderRepository) Create(order *SaleOrder) error {
	return r.DB.Create(order).Error
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
	return r.DB.Save(order).Error
}

func (r *SaleOrderRepository) Delete(id string) error {
	return r.DB.Delete(&SaleOrder{}, "id = ?", id).Error
}
