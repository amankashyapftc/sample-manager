package config

import (
	"fmt"
	"log"
	"sample-manager/model"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

var (
	db *gorm.DB
)

func DatabaseConnection() {
	host := "localhost"
	port := "5432"
	dbname := "sampleManagerDB"
	dbuser := "postgres"
	password := "12345"
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		dbuser,
		dbname,
		password,
	)

	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	d.AutoMigrate(&model.SampleItem{})

	if err != nil {
		log.Fatal(err)
	}

	db = d

	fmt.Println("Database connection successful...")
}

func GetDB() *gorm.DB {
	return db
}
