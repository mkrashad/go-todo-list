package api

import (
	"main_module/controllers"

	"github.com/gin-gonic/gin"
)

func Router() {
	r := gin.Default()
	// Tasks
	r.POST("/tasks", controllers.CreateTask)
	r.PUT("/tasks/:id", controllers.UpdateTask)
	r.GET("/tasks", controllers.GetAllTasks)
	r.GET("/tasks/:id", controllers.GetTaskById)
	r.DELETE("/tasks/:id", controllers.DeleteTask)

	// Users
	r.POST("/users", controllers.CreateUser)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.GET("/users", controllers.GetAllUsers)
	r.GET("/users/:id", controllers.GetUserById)
	r.DELETE("/users/:id", controllers.DeleteUser)
	r.Run()
}
