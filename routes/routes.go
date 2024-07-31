package routes

import (
	//"crud-api/controllers"

	"crud-api/controllers"

	"github.com/labstack/echo/v4"
)

func Routes() (e *echo.Echo) {
	e = echo.New()

	categoryRequest := e.Group("/category") // panggil postman
	categoryRequest.GET("", controllers.ReadAllCategories)
	categoryRequest.POST("/create", controllers.CreateCategory)
	categoryRequest.GET("/:id", controllers.ReadDetailCategories) // /:id namanya parameter
	categoryRequest.DELETE("/:id", controllers.DeleteCategory)    // /:id namanya parameter

	
	productRequest := e.Group("/product") // panggil postman
	productRequest.GET("", controllers.ReadAllProducts)
	productRequest.POST("/create", controllers.CreateProduct)
	productRequest.GET("/:id", controllers.ReadDetailProducts) // /:id namanya parameter
	productRequest.DELETE("/:id", controllers.DeleteProduct)   // /:id namanya parameter

	return
}
