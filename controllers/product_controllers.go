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

func CreateProduct(c *gin.Context) {
    var newProduct models.Product

    if err := c.ShouldBindJSON(&newProduct); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    newProduct.ID = nextID
    nextID++
    products = append(products, newProduct)
    c.JSON(http.StatusCreated, newProduct)
}

func UpdateProduct(c *gin.Context) {
    id := c.Param("id")
    for i, product := range products {
        if productID := c.Param("id"); productID == id {
            if err := c.ShouldBindJSON(&products[i]); err != nil {
                c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
                return
            }
            c.JSON(http.StatusOK, products[i])
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
}