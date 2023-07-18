package main

import (
	"main_module/controllers"
	initilalizers "main_module/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initilalizers.LoadEnvVariables()
	initilalizers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/tasks", controllers.AddTask)
	r.PUT("/tasks/:id", controllers.UpdateTask)
	r.GET("/tasks", controllers.GetAllTasks)
	r.GET("/tasks/:id", controllers.GetSingleTask)
	r.DELETE("/tasks/:id", controllers.DeleteTask)
	r.Run()
}
