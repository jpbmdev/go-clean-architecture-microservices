package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBconnection *gorm.DB

func ConnectToDB() {
	dns := "host=localhost user=username password=password dbname=customer port=5432"

	var err error
	DBconnection, err = gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	} else {
		fmt.Print("Database connected!")
	}
}
