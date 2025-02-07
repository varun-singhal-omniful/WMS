package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/varun-singhal-omniful/wms-service/models"
	"github.com/varun-singhal-omniful/wms-service/services"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateHub(c *gin.Context) {
	var hub models.Hub

	// Bind JSON request body to hub struct
	if err := c.ShouldBindJSON(&hub); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	// Call service layer
	err := services.CreateHub(&hub)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create hub"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Hub created successfully"})
}

// GetAllHubs handles fetching all hubs
func GetAllHubs(c *gin.Context) {
	hubs, err := services.GetAllHubs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve hubs"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"hubs": hubs})
}

// GetHub handles fetching a specific hub by ID
func GetHub(c *gin.Context) {
	id := c.Param("id")
	hubID, err := strconv.Atoi(id)
	// objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid hub ID"})
		return
	}
	hub, err := services.GetHub(hubID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Hub not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"hub": hub})
}

// // UpdateHub handles updating a specific hub by ID
// func UpdateHub(c *gin.Context) {
// 	id := c.Param("id")
// 	hubID, err := strconv.Atoi(id)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid hub ID"})
// 		return
// 	}

// 	var updatedHub models.Hub
// 	if err := c.ShouldBindJSON(&updatedHub); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
// 		return
// 	}

// 	hub, err := services.UpdateHub(c, hubID, updatedHub)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update hub"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Hub updated successfully", "hub": hub})
// }

// // DeleteHub handles deleting a specific hub by ID
// func DeleteHub(c *gin.Context) {
// 	id := c.Param("id")
// 	hubID, err := strconv.Atoi(id)
// 	// objID, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid hub ID"})
// 		return
// 	}

// 	err = services.DeleteHub(hubID)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete hub"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Hub deleted successfully"})
// }
