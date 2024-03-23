package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/ridhoafwani/fga-final-project/database"
	"github.com/ridhoafwani/fga-final-project/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Failed to load .env file")
	}
}

var (
	host     string
	port     string
	user     string
	password string
	dbName   string
)

func setData() {
	host = os.Getenv("DB_HOST")
	port = os.Getenv("DB_PORT")
	user = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	dbName = os.Getenv("DB_NAME")
}

func main() {
	dbName := os.Getenv("DB_NAME")

	rootDb := RootDatabaseConnection()

	var exists bool
	rootDb.Raw("SELECT EXISTS(SELECT FROM pg_database WHERE datname = ?)", dbName).Scan(&exists)

	if !exists {
		rootDb.Exec("CREATE DATABASE " + dbName)
	}

	db := database.DatabaseConnection()

	db.AutoMigrate(&models.User{}, &models.Photo{}, &models.Comment{}, &models.SocialMedia{})

	fmt.Println("Migration complete")

}

func RootDatabaseConnection() (db *gorm.DB) {
	setData()
	sqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, "postgres")

	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	fmt.Println("? Connected Successfully to the Database")
	return db
}
