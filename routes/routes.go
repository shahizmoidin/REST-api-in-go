package routes

import (
    "restapiingo/controllers"
    "github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

    router.GET("/products", controllers.GetProducts)
    router.GET("/products/:id", controllers.GetProduct)
    router.POST("/products", controllers.CreateProduct)
    router.PUT("/products/:id", controllers.UpdateProduct)
    router.DELETE("/products/:id", controllers.DeleteProduct)
	
}