package services

import (
	"github.com/Danangoffic/go-merce/app/controllers/request"
	"github.com/Danangoffic/go-merce/app/models"
	"github.com/Danangoffic/go-merce/app/repositories"
)

type service struct {
	repositories repositories.Repository
}

func InitService(repository repositories.Repository) *service {
	return &service{repositories: repository}
}

type Services interface {
	GetProducts(cateogry string) ([]models.Product, error)
	GetProductByID(ID int) (models.Product, error)
	CreateProduct(input request.StoreProduct) (models.Product, error)

	GetUserByID(ID uint) (models.User, error)

	GetUserTokenData(token string) (models.User, error)
}
