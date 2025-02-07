package services

import (
	"github.com/varun-singhal-omniful/wms-service/models"
	"github.com/varun-singhal-omniful/wms-service/repository"
)

// CreateInventory adds a new inventory record
func CreateInventory(inventory *models.Inventory) error {
	return repository.CreateInventory(inventory)
}

// GetAllInventory retrieves all inventory records
func GetAllInventory() ([]models.Inventory, error) {
	return repository.GetAllInventory()
}

// GetInventory retrieves a single inventory record by ID
func GetInventory(id int) (models.Inventory, error) {
	return repository.GetInventory(id)
}

func UpdateInventory(inventory *models.Inventory) error {
	existingInventory, err := repository.GetInventoryByHubAndSku(inventory.HubID, inventory.SkuID)
	if err != nil {
		// If inventory doesn't exist, create a new one
		return repository.CreateInventory(inventory)
	}

	// If inventory exists, update the quantity
	existingInventory.Quantity = inventory.Quantity
	return repository.UpdateInventory(&existingInventory)
}
