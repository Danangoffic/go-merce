package main

import (
	"fmt"

	"github.com/Danangoffic/go-merce/app/config"
	"github.com/Danangoffic/go-merce/app/database"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func migrate() {
	fmt.Printf("dsn = %v", config.DSN)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: config.DSN,
	}), &gorm.Config{})
	if err != nil {
		panic("failed connection with " + err.Error())
	}
	database.Migrate(db)
}
