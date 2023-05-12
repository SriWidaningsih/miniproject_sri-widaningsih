package database

import (
	"miniproject/config"
	"miniproject/models"
)

func GetOrders() (orders []models.OrderDetail, err error) {
	if err = config.DB.Find(&orders).Error; err != nil {
		return
	}
	return
}

func GetOrder(id uint) (order models.OrderDetail, err error) {
	order.ID = id
	if err = config.DB.First(&order).Error; err != nil {
		return
	}
	return
}

func UpdateOrder(order *models.OrderDetail) error {
	if err := config.DB.Updates(order).Error; err != nil {
		return err

	}
	return nil
}
func DeleteOrder(order *models.OrderDetail) error {
	if err := config.DB.Delete(order).Error; err != nil {
		return err

	}
	return nil
}

func CreateOrderDetail(order *models.OrderDetail) error {
	if arr := config.DB.Create(order).Error; arr != nil {
		return arr

	}
	return nil
}
