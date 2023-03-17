package main

import (
	"database/sql"
	"fmt"

	_ "github.com/godror/godror"
)

func oracle() {
	// set informasi koneksi
	host := "simasbanking.oracle.database"
	port := 1521
	sid := "ORCL"
	username := "SIMASBANKINGPROD"
	password := "simasbanking"
	// dialect := "org.hibernate.dialect.Oracle9Dialect"

	// buat string koneksi
	connectionString := fmt.Sprintf("%s/%s@//%s:%d/%s", username, password, host, port, sid)
	fmt.Printf("connection string with : %v", connectionString)

	// buat koneksi ke database
	db, err := sql.Open("godror", connectionString)
	if err != nil {
		panic(err.Error())
	}

	// test koneksi
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Connected to Oracle database!")
}
