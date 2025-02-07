package services

import (
	"github.com/varun-singhal-omniful/wms-service/models"
	"github.com/varun-singhal-omniful/wms-service/repository"
)

// CreateSKU calls the repository function to create an SKU
func CreateSKU(sku *models.Sku) error {
	return repository.CreateSKU(sku)
}

// GetAllSKUs fetches all SKUs from the repository
func GetAllSKUs() ([]models.Sku, error) {
	return repository.GetAllSKUs()
}

// GetSKU fetches a specific SKU by ID from the repository
func GetSKU(id int) (models.Sku, error) {
	return repository.GetSKU(id)
}
