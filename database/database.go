package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

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

func DatabaseConnection() (db *gorm.DB) {
	setData()
	sqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	fmt.Println("? Connected Successfully to the Database")
	return db
}
