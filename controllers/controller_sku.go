package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/varun-singhal-omniful/wms-service/models"
	"github.com/varun-singhal-omniful/wms-service/services"
)

// CreateSKU handles creating a new SKU
func CreateSKU(c *gin.Context) {
	var sku models.Sku

	// Bind JSON request body to SKU struct
	if err := c.ShouldBindJSON(&sku); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	// Call service layer
	err := services.CreateSKU(&sku)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create SKU"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "SKU created successfully"})
}

// GetAllSKUs handles fetching all SKUs
func GetAllSKUs(c *gin.Context) {
	skus, err := services.GetAllSKUs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve SKUs"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"skus": skus})
}

// GetSKU handles fetching a specific SKU by ID
func GetSKU(c *gin.Context) {
	id := c.Param("id")
	skuID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid SKU ID"})
		return
	}
	sku, err := services.GetSKU(skuID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "SKU not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"sku": sku})
}
