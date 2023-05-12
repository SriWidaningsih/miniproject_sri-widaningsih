package usecase

import (
	"errors"
	"miniproject/lib/database"
	"miniproject/models"
)

func GetListProducts() (products []models.Product, err error) {
	products, err = database.GetProducts()
	return
}

func GetProduct(id uint) (product models.Product, err error) {
	product, err = database.GetProduct(id)
	return
}

func UpdateProduct(product *models.Product) error {
	err := database.UpdateProduct(product)
	if err != nil {
		return err
	}
	return nil
}

func DeleteProduct(id uint) error {
	product := models.Product{}
	product.ID = id
	err := database.DeleteProduct(&product)
	if err != nil {
		return err
	}
	return nil
}

func CreateProduct(product *models.Product) error {
	if product.Name == "" {
		return errors.New("product name cannot be empty")
	}
	if product.Price == 0 {
		return errors.New("Product price is required")
	}
	if product.Price <= 0 {
		return errors.New("Product price must be greater than zero")
	}

	err := database.CreateProduct(product)
	if err != nil {
		return err
	}
	return nil
}
