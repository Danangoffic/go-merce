package seeders

import (
	"github.com/Danangoffic/go-merce/app/models"
	"github.com/bxcodec/faker/v3"
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

	for i := 0; i < 100; i++ {
		productName := faker.Name()
		slug := slug.Make(productName)
		price, _ := faker.RandomInt(50000, 1000000)
		qty, _ := faker.RandomInt(1, 500)
		userBy, _ := faker.RandomInt(1, 2)
		product := models.Product{
			Name:        productName,
			Slug:        slug,
			Type:        faker.Gender(),
			Description: faker.Paragraph(),
			Price:       price[0],
			Quantity:    qty[0],
			CreatedByID: uint(userBy[0]),
		}
		db.Create(&product)

		for j := 0; j < 5; j++ {
			productImages := models.ProductImages{
				ProductID:  product.ID,
				Url:        "https://images.pexels.com/photos/1152077/pexels-photo-1152077.jpeg?auto=compress&cs=tinysrgb&w=1600",
				IsFeatured: false,
			}
			db.Create(&productImages)
		}
	}

	// productName1 := string("Kaos Lengan Panjang")
	// slug1 := slug.Make(productName1)
	// product1 := models.Product{
	// 	Name:        productName1,
	// 	CreatedBy:   user1,
	// 	Slug:        slug1,
	// 	Type:        "Kaos",
	// 	Description: "Kaos lengan Panjang Pria yang cocok dipakai setiap saat",
	// 	Price:       99000,
	// 	Quantity:    25,
	// }
	// productName2 := string("Kaos Lengan Pendek")
	// slug2 := slug.Make(productName1)
	// product2 := models.Product{
	// 	Name:        productName2,
	// 	CreatedBy:   user2,
	// 	Slug:        slug2,
	// 	Type:        "Kaos",
	// 	Description: "Kaos lengan pendek Pria yang cocok dipakai setiap saat",
	// 	Price:       99000,
	// 	Quantity:    25,
	// }
	// db.Create(&product1).Create(&product2)

	// productImages := []models.ProductImages{{
	// 	ProductID:  product1.ID,
	// 	IsFeatured: false,
	// 	Url:        "https://images.pexels.com/photos/1152077/pexels-photo-1152077.jpeg?auto=compress&cs=tinysrgb&w=1600",
	// }, {
	// 	ProductID:  product1.ID,
	// 	IsFeatured: true,
	// 	Url:        "https://images.pexels.com/photos/2081199/pexels-photo-2081199.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1",
	// }, {
	// 	ProductID:  product2.ID,
	// 	IsFeatured: true,
	// 	Url:        "https://images.pexels.com/photos/984619/pexels-photo-984619.jpeg?auto=compress&cs=tinysrgb&w=1600",
	// }, {
	// 	ProductID:  product2.ID,
	// 	IsFeatured: false,
	// 	Url:        "https://images.pexels.com/photos/1687719/pexels-photo-1687719.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1",
	// }}
	// db.Create(&productImages)
}
