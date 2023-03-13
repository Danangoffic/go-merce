package seeders

import (
	"github.com/Danangoffic/go-merce/app/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RunSeederUser(db *gorm.DB) {
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
}
