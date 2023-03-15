package repositories

import (
	"errors"
	"fmt"

	"github.com/Danangoffic/go-merce/app/models"
)

func (p *RepositoryApp) GetProducts(category string) ([]models.Product, error) {
	var products []models.Product
	query := p.db.Preload("CreatedBy").Preload("ProductImages").Order("ID DESC")
	if category != "" {
		fmt.Printf("category ada : %v", category)
		query.Where("type = ?", category)
	}
	fmt.Printf("query generated : %v\n", query)
	if err := query.Find(&products).Error; err != nil {
		return nil, err
	}

	if len(products) == 0 {
		return nil, errors.New("Products data not found")
	}
	return products, nil
}

func (p *RepositoryApp) GetProduct(slug string) (models.Product, error) {
	var product models.Product
	if err := p.db.Where(models.Product{Slug: slug}).Find(&product).Error; err != nil {
		return product, errors.New("Failed to search")
	}
	return product, nil
}

func (r *RepositoryApp) CreateProduct(product models.Product) (models.Product, error) {
	if err := r.db.Create(&product).Error; err != nil {
		return models.Product{}, err
	}
	return product, nil
}

func (r *RepositoryApp) GetProductByID(ID int) (models.Product, error) {
	var product models.Product
	if err := r.db.Where(models.Product{ID: uint(ID)}).Find(&product).Error; err != nil {
		return product, err
	}
	return product, nil
}
