package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/korisnik/ecomm-backend/config"
	"github.com/korisnik/ecomm-backend/models"
)

// GetProducts - vraća sve proizvode
func GetProducts(c *gin.Context) {
	var products []models.Product
	config.DB.Find(&products)

	c.JSON(http.StatusOK, products)
}

// GetProductByID - vraća proizvod na osnovu ID
func GetProductByID(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Proizvod nije pronađen"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// CreateProduct - kreira novi proizvod
func CreateProduct(c *gin.Context) {
	var product models.Product

	// Učitavamo JSON podatke u product struct
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&product)
	c.JSON(http.StatusCreated, product)
}

// UpdateProduct - ažurira postojeći proizvod
func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	// Prvo pronalazimo proizvod po ID
	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Proizvod nije pronađen"})
		return
	}

	// Učitavamo nove vrednosti iz JSON-a
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&product)
	c.JSON(http.StatusOK, product)
}

// DeleteProduct - briše postojeći proizvod
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	// Proveravamo da li proizvod postoji
	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Proizvod nije pronađen"})
		return
	}

	config.DB.Delete(&product)
	c.JSON(http.StatusOK, gin.H{"message": "Proizvod je obrisan"})
}
