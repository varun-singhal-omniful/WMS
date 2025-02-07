package services

import (
	"github.com/varun-singhal-omniful/wms-service/models"
	"github.com/varun-singhal-omniful/wms-service/repository"
)

// CreateTenant creates a new tenant
func CreateTenant(tenant *models.Tenant) error {
	return repository.CreateTenant(tenant)
}

// GetAllTenants fetches all tenants
func GetAllTenants() ([]models.Tenant, error) {
	return repository.GetAllTenants()
}

// GetTenant fetches a specific tenant by ID
func GetTenant(id int) (models.Tenant, error) {
	return repository.GetTenant(id)
}
