package routes

import (
	"restapiingo/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)

	secured := router.Group("/").Use(controllers.Authenticate())
	{
		secured.GET("/products", controllers.GetProducts)
		secured.GET("/products/:id", controllers.GetProduct)
		secured.POST("/products", controllers.CreateProduct)
		secured.PUT("/products/:id", controllers.UpdateProduct)
		secured.DELETE("/products/:id", controllers.DeleteProduct)
	}
}
