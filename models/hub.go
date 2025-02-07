package models

type Hub struct {
	ID             int    `gorm:"primaryKey" json:"id"`
	TenantID       int    `json:"tenant_id"`
	ManagerName    string `json:"manager_name"`
	ManagerContact string `json:"manager_contact"`
	ManagerEmail   string `json:"manager_email"`
}
