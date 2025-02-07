package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/varun-singhal-omniful/wms-service/models"
	"github.com/varun-singhal-omniful/wms-service/services"
)

// CreateInventory handles the creation of an inventory entry
func CreateInventory(c *gin.Context) {
	var inventory models.Inventory

	// Bind JSON request body to inventory struct
	if err := c.ShouldBindJSON(&inventory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	// Call service layer
	err := services.CreateInventory(&inventory)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create inventory"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Inventory created successfully"})
}

// GetAllInventory handles fetching all inventory records
func GetAllInventory(c *gin.Context) {
	inventories, err := services.GetAllInventory()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve inventory"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"inventory": inventories})
}

// GetInventory handles fetching a specific inventory item by ID
func GetInventory(c *gin.Context) {
	id := c.Param("id")
	inventoryID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid inventory ID"})
		return
	}
	inventory, err := services.GetInventory(inventoryID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Inventory item not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"inventory": inventory})
}
func UpdateInventory(c *gin.Context) {
	var inventory models.Inventory

	// Bind JSON request body to inventory struct
	if err := c.ShouldBindJSON(&inventory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	// Call service layer to update or create inventory
	err := services.UpdateInventory(&inventory)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update inventory"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Inventory updated successfully"})
}
