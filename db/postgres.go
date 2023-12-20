package db

// docker run --name my-postgres -e POSTGRES_PASSWORD=root -p 8081:5432 -d postgres

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=root dbname=postgres port=8180 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	DB = database

	return DB
}

func CloseDB() {
	if DB != nil {
		sqlDB, _ := DB.DB()
		sqlDB.Close()
	}
}
