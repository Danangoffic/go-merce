package repositories

import (
	"github.com/Danangoffic/go-merce/app/models"
	"gorm.io/gorm"
)

type RepositoryApp struct {
	db *gorm.DB
}

// To Init Repository interface
func GetRepository(db *gorm.DB) Repository {
	return &RepositoryApp{
		db: db,
	}
}

type Repository interface {
	// Product Interfaces
	GetProducts(category string) ([]models.Product, error)
	GetProduct(slug string) (models.Product, error)
	GetProductByID(ID int) (models.Product, error)
	CreateProduct(product models.Product) (models.Product, error)

	// User Interfaces
	GetUserByID(user models.User) (models.User, error)
	CreateUser(user models.User) (models.User, error)
}
