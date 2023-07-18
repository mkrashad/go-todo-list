package main

import (
	"fmt"
	initilalizers "main_module/initializers"
	"main_module/models"
)

func init() {
	initilalizers.LoadEnvVariables()
	initilalizers.ConnectToDB()
}

func main() {
	initilalizers.DB.AutoMigrate(&models.Task{}, &models.User{})
	fmt.Println("Database migrated succesfully")
}
