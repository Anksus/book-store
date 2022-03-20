package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func connect() {
	dbURL := "postgres://postgres:postgres@localhost:5432/bookstore-postgres"

	d, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db = d
}

func getDB() *gorm.DB {
	return db
}
