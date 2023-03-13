package request

import "github.com/Danangoffic/go-merce/app/models"

type StoreProduct struct {
	CreatedByID string      `json:"created_by_id" binding:"required"`
	CreatedBy   models.User `gorm:"foreignKey(CreatedByID)" json:"created_by"`
	Name        string      `json:"name" binding:"required"`
	Slug        string
	Type        string `json:"type" binding:"required"`
	Description string `json:"description"`
	Price       int    `json:"price" binding:"required"`
	Quantity    int    `json:"quantity" binding:"required"`
}
