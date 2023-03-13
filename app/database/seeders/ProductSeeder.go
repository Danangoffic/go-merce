package seeders

import (
	"github.com/Danangoffic/go-merce/app/models"
	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

func RunSeederProduct(db *gorm.DB) {
	var user1 models.User
	var user2 models.User
	user1.ID = 1
	user2.ID = 2
	db.Find(&user1)
	db.Find(&user2)

	productName1 := string("Kaos Lengan Panjang")
	slug1 := slug.Make(productName1)
	product1 := models.Product{
		Name:        productName1,
		CreatedBy:   user1,
		Slug:        slug1,
		Type:        "Kaos",
		Description: "Kaos lengan Panjang Pria yang cocok dipakai setiap saat",
		Price:       99000,
		Quantity:    25,
	}
	productName2 := string("Kaos Lengan Pendek")
	slug2 := slug.Make(productName1)
	product2 := models.Product{
		Name:        productName2,
		CreatedBy:   user2,
		Slug:        slug2,
		Type:        "Kaos",
		Description: "Kaos lengan pendek Pria yang cocok dipakai setiap saat",
		Price:       99000,
		Quantity:    25,
	}
	db.Create(&product1).Create(&product2)
}
