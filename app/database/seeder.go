package database

import (
	"github.com/Danangoffic/go-merce/app/models"
	"github.com/gosimple/slug"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedDB(db *gorm.DB) {
	if db.Migrator().HasTable(&models.User{}) && db.Migrator().HasTable(&models.Product{}) {
		db.Migrator().DropTable(&models.User{}, &models.Product{})
	}
	db.AutoMigrate(&models.User{}, &models.Product{})
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("123123"), bcrypt.DefaultCost)
	user1 := models.User{
		Name:     "Danang Arif",
		Email:    "darifrahmanda@gmail.com",
		Password: string(hashedPassword),
	}
	user2 := models.User{
		Name:     "Riselda Rahma",
		Email:    "risel123@gmail.com",
		Password: string(hashedPassword),
	}
	db.Create(&user1).Create(&user2)

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
