package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// GetDBConfig retrieves the database configuration from environment variables
func GetDBConfig() (string, string, string, string, string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbName := os.Getenv("DBNAME")
	dbUser := os.Getenv("DBUSER")
	dbPass := os.Getenv("DBPASS")
	dbHost := os.Getenv("DBHOST")
	dbPort := os.Getenv("DBPORT")
	return dbName, dbUser, dbPass, dbHost, dbPort
}

// DBInitialize initializes the database connection using the configuration from environment variables
func DBInitialize() *gorm.DB {
	dbName, dbUser, dbPass, dbHost, dbPort := GetDBConfig()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbHost, dbUser, dbPass, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to connect to database: %s", err))
	}
	return db

}
