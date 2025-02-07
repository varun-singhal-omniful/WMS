package router

import (
	"context"
	// "net/http"
	"github.com/omniful/go_commons/http"
	"github.com/varun-singhal-omniful/wms-service/controllers"
)

func Initialize(ctx context.Context, s *http.Server) (err error) {
	// skus
	SkuV1 := s.Engine.Group("/api/v1/sku")
	SkuV1.POST("/", controllers.CreateSKU)
	SkuV1.GET("/", controllers.GetAllSKUs)
	SkuV1.GET("/:id", controllers.GetSKU)

	// inventory
	InventoryV1 := s.Engine.Group("/api/v1/inventory")
	InventoryV1.POST("/", controllers.CreateInventory)
	InventoryV1.GET("/", controllers.GetAllInventory)
	InventoryV1.GET("/:id", controllers.GetInventory)
	InventoryV1.PUT("/:id", controllers.GetInventory)

	// Hub
	HubV1 := s.Engine.Group("/api/v1/hub")
	HubV1.POST("/", controllers.CreateHub)
	HubV1.GET("/", controllers.GetAllHubs)
	HubV1.GET("/:id", controllers.GetHub)

	// Tenant
	TenantV1 := s.Engine.Group("/api/v1/tenant")
	TenantV1.POST("/", controllers.CreateTenant)
	TenantV1.GET("/", controllers.GetAllTenants)
	TenantV1.GET("/:id", controllers.GetTenant)

	return
}
