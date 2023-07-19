package main

import (
	"fmt"
	"log"
	initilalizers "main_module/initializers"
	"main_module/models"
)

func init() {
	initilalizers.LoadEnvVariables()
	initilalizers.ConnectToDB()
}

func main() {
	err := initilalizers.DB.AutoMigrate(&models.Task{}, &models.User{})
	if err != nil {
		log.Fatal("Could not migrate:", err)
	}
	fmt.Println("Database migrated succesfully")
}
