package controllers

import (
    "net/http"
    "restapiingo/models"
    "github.com/gin-gonic/gin"
)

var products []models.Product
var nextID = 1


func GetProducts(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, products)
}


func GetProduct(c *gin.Context) {
    id := c.Param("id")
    for _, product := range products {
        if productID := c.Param("id"); productID == id {
            c.IndentedJSON(http.StatusOK, product)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Product not found"})
}


func CreateProduct(c *gin.Context) {
    var newProduct models.Product

    if err := c.ShouldBindJSON(&newProduct); err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    newProduct.ID = uint(nextID)
    nextID++
    products = append(products, newProduct)
    c.IndentedJSON(http.StatusCreated, newProduct)
}


func UpdateProduct(c *gin.Context) {
    id := c.Param("id")
    for i := range products {
        if productID := c.Param("id"); productID == id {
            if err := c.ShouldBindJSON(&products[i]); err != nil {
                c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
                return
            }
            c.IndentedJSON(http.StatusOK, products[i])
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Product not found"})
}


func DeleteProduct(c *gin.Context) {
    id := c.Param("id")
    for i := range products {
        if productID := c.Param("id"); productID == id {
            products = append(products[:i], products[i+1:]...)
            c.IndentedJSON(http.StatusOK, gin.H{"message": "Product deleted"})
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Product not found"})
}
