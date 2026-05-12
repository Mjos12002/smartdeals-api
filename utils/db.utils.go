package utils

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// GetDBConfig retrieves the database configuration from environment variables
func GetDBConfig() (string, string, string, string, string) {
	dbName := "smartdeals"
	dbUser := "smartdeals"
	dbPass := "4d135b886ad1527499bb73c7"
	dbHost := "localhost"
	return dbName, dbUser, dbPass, dbHost, "5432"
}

// DBInitialize initializes the database connection using the configuration from environment variables
func DBInitialize() *gorm.DB {
	dbName, dbUser, dbPass, dbHost, dbPort := GetDBConfig()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPass, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	return db

}
