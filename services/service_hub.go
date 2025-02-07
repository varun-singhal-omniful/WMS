package services

import (
	"github.com/varun-singhal-omniful/wms-service/models"
	"github.com/varun-singhal-omniful/wms-service/repository"
)

func CreateHub(hub *models.Hub) error {
	return repository.CreateHub(hub)

}

func GetAllHubs() ([]models.Hub, error) {
	return repository.GetAllHubs()
}

func GetHub(id int) (models.Hub, error) {
	return repository.GetHub(id)
}

// func UpdateHub(id int64) {
// 	return repository.UpdateHub(hub)

// }

// func DeleteHub(id int)error {
// 	return repository.DeleteHub(id)

// }
