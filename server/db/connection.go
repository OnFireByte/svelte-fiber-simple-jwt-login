package db

import (
	"fmt"
	"os"

	"github.com/onfirebyte/simple-jwt-login/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {
	dbData := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_DBNAME"),
		os.Getenv("POSTGRES_PASSWORD"))
	config := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
 	}
	connection, err := gorm.Open(postgres.Open(dbData), config)
	if err != nil {
		panic(err)
	}

	DB = connection

	connection.AutoMigrate(&models.User{},&models.Note{})

}