package main

import (
	"fmt"
	"time"

	"github.com/Danangoffic/go-merce/app/config"
	"github.com/Danangoffic/go-merce/app/database"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func seeder() {
	db, _ := gorm.Open(mysql.New(mysql.Config{
		DSN: config.DSN,
	}), &gorm.Config{})
	timer := time.Now()
	fmt.Printf("current time : %v\n", timer)
	database.SeedDB(db)
	fmt.Printf("\nFinished at : %v\n", time.Now())
}
