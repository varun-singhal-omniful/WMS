package repository

import (
	"github.com/varun-singhal-omniful/wms-service/appinit"
	"github.com/varun-singhal-omniful/wms-service/models"
)

// CreateTenant inserts a new tenant record into the database
func CreateTenant(tenant *models.Tenant) error {
	return appinit.DB.Create(tenant).Error
}

// GetAllTenants fetches all tenant records from the database
func GetAllTenants() ([]models.Tenant, error) {
	var tenants []models.Tenant
	if err := appinit.DB.Find(&tenants).Error; err != nil {
		return nil, err
	}
	return tenants, nil
}

// GetTenant fetches a specific tenant record by ID
func GetTenant(id int) (models.Tenant, error) {
	var tenant models.Tenant
	if err := appinit.DB.First(&tenant, id).Error; err != nil {
		return tenant, err
	}
	return tenant, nil
}
