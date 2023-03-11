package product

import (
	"github.com/Danangoffic/go-merce/app/models"
	"github.com/Danangoffic/go-merce/app/repositories/product"
)

type productService struct {
	productRepo product.Repository
}

func InitService(productRepo product.Repository) Services {
	return &productService{productRepo: productRepo}
}

type Services interface {
	GetProducts(cateogry string) ([]models.Product, error)
}
