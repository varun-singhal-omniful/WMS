package repository

import (
	"github.com/varun-singhal-omniful/wms-service/appinit"
	"github.com/varun-singhal-omniful/wms-service/models"
)

// CreateInventory adds a new inventory record
func CreateInventory(inventory *models.Inventory) error {
	return appinit.DB.Create(inventory).Error
}

// GetAllInventory retrieves all inventory records
func GetAllInventory() ([]models.Inventory, error) {
	var inventories []models.Inventory
	if err := appinit.DB.Find(&inventories).Error; err != nil {
		return nil, err
	}
	return inventories, nil
}

// GetInventory retrieves a single inventory record by ID
func GetInventory(id int) (models.Inventory, error) {
	var inventory models.Inventory
	if err := appinit.DB.First(&inventory, id).Error; err != nil {
		return inventory, err
	}
	return inventory, nil
}
func GetInventoryByHubAndSku(hubID int, skuID int) (models.Inventory, error) {
	var inventory models.Inventory
	err := appinit.DB.Where("hub_id = ? AND sku_id = ?", hubID, skuID).First(&inventory).Error
	return inventory, err
}

func UpdateInventory(inventory *models.Inventory) error {
	return appinit.DB.Save(inventory).Error
}
