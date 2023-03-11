package product

import "github.com/Danangoffic/go-merce/app/models"

func (p *productService) GetProducts(category string) ([]models.Product, error) {
	return p.productRepo.GetProducts(category)
}
