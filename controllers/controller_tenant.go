package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/varun-singhal-omniful/wms-service/models"
	"github.com/varun-singhal-omniful/wms-service/services"
)

func CreateTenant(c *gin.Context) {
	var tenant models.Tenant

	// Bind JSON request body to tenant struct
	if err := c.ShouldBindJSON(&tenant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	// Call service layer
	err := services.CreateTenant(&tenant)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create tenant"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Tenant created successfully"})
}

// GetAllTenants handles fetching all tenants
func GetAllTenants(c *gin.Context) {
	tenants, err := services.GetAllTenants()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tenants"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"tenants": tenants})
}

// GetTenant handles fetching a specific tenant by ID
func GetTenant(c *gin.Context) {
	id := c.Param("id")
	tenantID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tenant ID"})
		return
	}
	tenant, err := services.GetTenant(tenantID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tenant not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"tenant": tenant})
}
