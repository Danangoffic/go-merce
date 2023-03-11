package product

import (
	"github.com/Danangoffic/go-merce/app/models"
	"gorm.io/gorm"
)

type productRepo struct {
	db *gorm.DB
}

func GetRepository(db *gorm.DB) Repository {
	return &productRepo{
		db: db,
	}
}

type Repository interface {
	GetProducts(category string) ([]models.Product, error)
	GetProduct(slug string) (models.Product, error)
}
