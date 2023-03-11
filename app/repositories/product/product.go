package product

import (
	"errors"

	"github.com/Danangoffic/go-merce/app/models"
)

func (p *productRepo) GetProducts(category string) ([]models.Product, error) {
	var products []models.Product
	if err := p.db.Find(&products).Order("ID DESC").Error; err != nil {
		return nil, err
	}

	if len(products) == 0 {
		return nil, errors.New("Products data not found")
	}
	return products, nil
}

func (p *productRepo) GetProduct(slug string) (models.Product, error) {
	var product models.Product
	if err := p.db.Where(models.Product{Slug: slug}).Find(&product).Error; err != nil {
		return product, errors.New("Failed to search")
	}
	return product, nil
}
