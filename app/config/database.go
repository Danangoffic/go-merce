package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// InitDB function will initialize the database, run the migration table, seed the database if needed
// and will return as gorm.DB pointer
func InitDB() *gorm.DB {
	// create a database connection
	db, err := setupDatabase()

	if err != nil {
		panic(err.Error())
	}
	sqlDB, err := db.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(50)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}

// setupDatabase function will create a connection to database, depends on constant.go, you can change the database configuration.
// This function will return gorm.DB pointer and connection error
func setupDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: DSN,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

// MongoConnection func will connect to mongo database using no-sql mongo driver
// and return the *mongo.client object to interacts with mongo database and collections
func MongoConnection() *mongo.Client {
	client, _ := connectWithMongoDB()
	return client
}

func connectWithMongoDB() (*mongo.Client, error) {
	fmt.Printf("string connection to : %v\n", MongoDBURL)

	// set server API options to mongodb
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(MongoDBURL).
		SetServerAPIOptions(serverAPIOptions)

	// set context logging
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// to cancel if meet timeout
	// defer cancel()

	// to connect mongo db with clientOptions
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
		panic("failed to connect with mongo db with : " + err.Error())
	}

	return client, nil
}
