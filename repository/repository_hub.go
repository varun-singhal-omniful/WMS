package repository

import (
	"github.com/varun-singhal-omniful/wms-service/appinit"
	"github.com/varun-singhal-omniful/wms-service/models"
)

func CreateHub(hub *models.Hub) error {
	return appinit.DB.Create(hub).Error
}
func GetAllHubs() ([]models.Hub, error) {
	var hubs []models.Hub
	if err := appinit.DB.Find(&hubs).Error; err != nil {
		return nil, err
	}
	return hubs, nil
}
func GetHub(id int) (models.Hub, error) {
	var hub models.Hub
	if err := appinit.DB.First(&hub, id).Error; err != nil {
		return hub, err
	}
	return hub, nil

}
