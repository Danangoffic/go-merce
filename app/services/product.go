package services

import (
	"github.com/Danangoffic/go-merce/app/controllers/request"
	"github.com/Danangoffic/go-merce/app/models"
	"github.com/gosimple/slug"
)

func (s *service) GetProducts(category string) ([]models.Product, error) {
	return s.repositories.GetProducts(category)
}

func (s *service) CreateProduct(input request.StoreProduct) (models.Product, error) {
	newproduct := models.Product{}
	newproduct.Name = input.Name
	newproduct.Slug = slug.Make(input.Name)
	newproduct.Price = input.Price
	newproduct.Type = input.Type
	newproduct.Description = input.Description
	newproduct.CreatedByID = input.CreatedBy.ID
	newproduct.Quantity = input.Quantity
	newproduct.Price = input.Price
	responCreate, err := s.repositories.CreateProduct(newproduct)
	if err != nil {
		return responCreate, err
	}
	return responCreate, nil
}
