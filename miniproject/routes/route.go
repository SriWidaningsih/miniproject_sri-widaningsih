package routes

import (
	"miniproject/controllers"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func New(e *echo.Echo, db *gorm.DB) {
	e.Validator = &CustomValidator{validator: validator.New()}

	e.POST("/login", controllers.LoginUserController)
	e.POST("/register", controllers.CreateUserController)

	// user collection
	user := e.Group("/users")
	user.GET("", controllers.GetUsersController)
	user.GET("/:id", controllers.GetUserController)
	user.POST("", controllers.CreateUserController)
	user.PUT("/:id", controllers.UpdateUserController)
	user.DELETE("/:id", controllers.DeleteUserController)

	//Product by pembeli
	product := e.Group("/product")
	product.GET("", controllers.GetProductsController)

	// Product by admin
	product.GET("/:id", controllers.GetProductController)
	product.POST("", controllers.CreateProductController)
	product.DELETE("/:id", controllers.DeleteProductController)
	product.PUT("/:id", controllers.UpdateProductController)
  
	//Order by Admin
	order := e.Group("/order")
	order.GET("", controllers.GetOrdersController)
	order.GET("/:id", controllers.GetOrderController)
	order.POST("", controllers.CreateOrderController)
	order.DELETE("/:id", controllers.DeleteOrderController)
	order.PUT("/:id", controllers.UpdateOrderController)

}
