package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/ridhoafwani/fga-final-project/database"
	"github.com/ridhoafwani/fga-final-project/models"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Failed to load .env file")
	}
	db := database.DatabaseConnection()

	db.AutoMigrate(&models.User{}, &models.Photo{}, &models.Comment{}, &models.SocialMedia{})

	fmt.Println("Migration complete")
}
