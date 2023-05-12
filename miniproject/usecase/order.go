package usecase

import (
	"errors"
	"miniproject/lib/database"
	"miniproject/models"
)

func GetListOrders() (orders []models.OrderDetail, err error) {
	orders, err = database.GetOrders()
	return
}

func GetOrder(id uint) (order models.OrderDetail, err error) {
	order, err = database.GetOrder(id)
	return
}

func UpdateOrder(order *models.OrderDetail) error {
	err := database.UpdateOrder(order)
	if err != nil {
		return err
	}
	return nil
}

func DeleteOrder(id uint) error {
	order := models.OrderDetail{}
	order.ID = id
	err := database.DeleteOrder(&order)
	if err != nil {
		return err
	}
	return nil
}

func CreateOrder(order *models.OrderDetail) error {
	if order.Quantity == 0 {
		return errors.New("quantity cannot be empty")
	}
	if order.UserID == 0 {
		return errors.New("user ID is required")
	}
	if order.ProductID == 0 {
		return errors.New("product ID is required")
	}

	err := database.CreateOrderDetail(order)
	if err != nil {
		return err
	}
	return nil
}
