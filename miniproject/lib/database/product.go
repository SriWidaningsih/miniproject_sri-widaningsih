package database

import (
	"miniproject/config"
	"miniproject/models"
)

func GetProducts() (products []models.Product, err error) {
	if err = config.DB.Find(&products).Error; err != nil {
		return
	}
	return
}

func GetProduct(id uint) (product models.Product, err error) {
	product.ID = id
	if err = config.DB.First(&product).Error; err != nil {
		return
	}
	return
}

func UpdateProduct(product *models.Product) error {
	if err := config.DB.Updates(product).Error; err != nil {
		return err

	}
	return nil
}
func DeleteProduct(product *models.Product) error {
	if err := config.DB.Delete(product).Error; err != nil {
		return err

	}
	return nil
}

func CreateProduct(product *models.Product) error {
	if arr := config.DB.Create(product).Error; arr != nil {
		return arr

	}
	return nil
}
