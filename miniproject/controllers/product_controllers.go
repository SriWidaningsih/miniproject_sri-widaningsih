package controllers

import (
	"net/http"
	"strconv"

	"miniproject/lib/database"
	"miniproject/models"
	"miniproject/usecase"

	"github.com/labstack/echo/v4"

	"miniprojek/config"
)

// get product by id
func GetProductController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	product, err := database.GetProduct(uint(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         " error get Product",
			"errorDescription": err,
		})

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  product,
	})
}

//create new product

func CreateProductController(c echo.Context) error {
	product := models.Product{}
	c.Bind(&product)

	if err := usecase.CreateProduct(&product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         " error create product",
			"errorDescription": err,
		})

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create new product",
		"users":    product,
	})
}

// delete product by id
func DeleteProductController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := usecase.DeleteProduct(uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error delete product",
			"errorDescription": err,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "succses delete product",
	})
}

// update product by id
func UpdateProductController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	product := models.Product{}
	c.Bind(&product)
	product.ID = uint(id)
	if err := usecase.UpdateProduct(&product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":          "error update product",
			"errorDescription": err,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "succsess update product",
	})
}

// get all products
func GetProductsController(c echo.Context) error {
	var products []models.Product

	if err := config.DB.Find(&products).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all products",
		"users":   products,
	})
}
