package main

import (
	"main_module/api"
	initilalizers "main_module/db"
)

func init() {
	initilalizers.LoadEnvVariables()
	initilalizers.ConnectToDB()
	initilalizers.SyncDB()
}

func main() {
	api.Router()
}
