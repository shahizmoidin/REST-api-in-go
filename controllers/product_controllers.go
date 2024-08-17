package controllers

import (
    "net/http"
    "restapiingo/models"
    "github.com/gin-gonic/gin"
)

var products []models.Product
var nextID = 1


func GetProducts(c *gin.Context) {
    c.JSON(http.StatusOK, products)
}

func GetProduct(c *gin.Context) {
    id := c.Param("id")
    for _, product := range products {
        if productID := c.Param("id"); productID == id {
            c.JSON(http.StatusOK, product)
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
}