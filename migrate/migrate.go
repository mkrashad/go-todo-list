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
	initilalizers.DB.AutoMigrate(&models.Task{})
	fmt.Println("Database migrated succesfully")
}
