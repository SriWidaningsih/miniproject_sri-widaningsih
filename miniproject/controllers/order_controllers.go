package controllers

import (
	"net/http"
	"strconv"

	"miniproject/lib/database"
	"miniproject/models"
	"miniproject/usecase"

	"github.com/labstack/echo/v4"

	"miniproject/config"
)

// get order by id
func GetOrderController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	order, err := database.GetOrder(uint(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         " error get Order",
			"errorDescription": err,
		})

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  order,
	})
}

//create new order

func CreateOrderController(c echo.Context) error {
	order := models.OrderDetail{}
	c.Bind(&order)

	if err := usecase.CreateOrder(&order); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         " error create order",
			"errorDescription": err,
		})

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create new order",
		"users":    order,
	})
}

// delete order by id
func DeleteOrderController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := usecase.DeleteOrder(uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error delete order",
			"errorDescription": err,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "succses delete order",
	})
}

// update order by id
func UpdateOrderController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	order := models.OrderDetail{}
	c.Bind(&order)
	order.ID = uint(id)
	if err := usecase.UpdateOrder(&order); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":          "error update order",
			"errorDescription": err,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "succsess update order",
	})
}

// get all orders
func GetOrdersController(c echo.Context) error {
	var orders []models.OrderDetail

	if err := config.DB.Find(&orders).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all orders",
		"users":   orders,
	})
}
