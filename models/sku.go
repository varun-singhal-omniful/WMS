package models

type Sku struct {
	ID          int    `gorm:"primaryKey" json:"id"`
	ProductID   int    `json:"product_id"`
	HubID       int    `json:"hub_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
