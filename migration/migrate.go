package main

import (
	"fmt"

	"github.com/ridhoafwani/fga-final-project/database"
	"github.com/ridhoafwani/fga-final-project/models"
)

func main() {
	dbName := "mygram"

	rootDb := database.RootDatabaseConnection()

	var exists bool
	rootDb.Raw("SELECT EXISTS(SELECT FROM pg_database WHERE datname = ?)", dbName).Scan(&exists)

	if !exists {
		rootDb.Exec("CREATE DATABASE " + dbName)
	}

	db := database.DatabaseConnection()

	db.AutoMigrate(&models.User{}, &models.Photo{}, &models.Comment{}, &models.SocialMedia{})
	fmt.Println("Migration complete")

}
