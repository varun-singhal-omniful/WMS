package models

type Inventory struct {
	ID       int `gorm:"primaryKey" json:"id"`
	HubID    int `json:"hub_id"`
	SkuID    int `json:"sku_id"`
	Quantity int `json:"quantity"`
}
