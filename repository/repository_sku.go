package repository

import (
	"github.com/varun-singhal-omniful/wms-service/appinit"
	"github.com/varun-singhal-omniful/wms-service/models"
)

// CreateSKU inserts a new SKU record into the database
func CreateSKU(sku *models.Sku) error {
	return appinit.DB.Create(sku).Error
}

// GetAllSKUs retrieves all SKU records from the database
func GetAllSKUs() ([]models.Sku, error) {
	var skus []models.Sku
	if err := appinit.DB.Find(&skus).Error; err != nil {
		return nil, err
	}
	return skus, nil
}

// GetSKU retrieves a specific SKU by ID from the database
func GetSKU(id int) (models.Sku, error) {
	var sku models.Sku
	if err := appinit.DB.First(&sku, id).Error; err != nil {
		return sku, err
	}
	return sku, nil
}
