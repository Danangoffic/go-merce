package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Danangoffic/go-merce/app/config"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	// get mongo db connection and its configs
	client := config.MongoConnection()
	ctx := context.TODO()

	// to get all database list
	_, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	// to get vouchers collection in v_commerce database
	VouchersCollection := client.Database("v_commerce").Collection("vouchers")
	newIntface := bson.D{{}}
	VouchersCollection.DeleteMany(ctx, newIntface)

	// sample data to store a data to mongo db using bson.D interfaces
	docs := []interface{}{
		bson.D{
			{"title", "PUBGM-VOUCHER-100"}, {"price", 100000}, {"type", "PUBG"}},
		bson.D{{"title", "PUBGM-VOUCHER-150"}, {"price", 150000}, {"type", "PUBG"}},
		bson.D{{"title", "PUBGM-VOUCHER-200"}, {"price", 200000}, {"type", "PUBG"}},
	}

	// sample to insert many data on vouchers collection in v_commerce database
	res, err := VouchersCollection.InsertMany(ctx, docs)
	if err != nil {
		fmt.Printf("Failed to insert data to vouchers collection with : %v\n", err.Error())
	}

	// Find all data in vouchers collection
	allVouchersCollection, err := VouchersCollection.Find(ctx, bson.D{{}})
	if err != nil {
		panic("error get all collection voucher with : " + err.Error())
	}

	// set collection data to fetchData mapper as []bson.M
	var fetchData []bson.M
	if err = allVouchersCollection.All(ctx, &fetchData); err != nil {
		panic("Failed to get all vouchers collection")
	}

	// iterate all fetchData that returned from fetchData
	for _, result := range fetchData {
		fmt.Printf("data fetched = %v\n", result)
	}

	fmt.Printf("data inserted in vouchers collection : %v\n", res)
}
