package main

import (
	"main_module/api"
	initilalizers "main_module/initializers"
)

func init() {
	initilalizers.LoadEnvVariables()
	initilalizers.ConnectToDB()
}

func main() {
	api.Router()
}
