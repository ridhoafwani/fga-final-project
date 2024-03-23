package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
)

func SetData() {
	Host = os.Getenv("DB_HOST")
	Port = os.Getenv("DB_PORT")
	User = os.Getenv("DB_USER")
	Password = os.Getenv("DB_PASSWORD")
	DbName = os.Getenv("DB_NAME")
}

func DatabaseConnection() (db *gorm.DB) {
	sqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", Host, Port, User, Password, DbName)

	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	fmt.Println("? Connected Successfully to the Database")
	return db
}
