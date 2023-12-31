package initilalizers

import (
	"fmt"
	"log"
	"main_module/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func ConnectToDB() {
	var err error
	DbConfig := struct {
		Host     string
		User     string
		Password string
		DbName   string
		Port     string
	}{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DbName:   os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", DbConfig.Host, DbConfig.User, DbConfig.Password, DbConfig.DbName, DbConfig.Port)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database!")
	}
}

func SyncDB() {
	err := DB.AutoMigrate(&models.Task{}, &models.User{})
	if err != nil {
		log.Fatal("Could not migrate:", err)
	}
	fmt.Println("Database migrated succesfully")
}
