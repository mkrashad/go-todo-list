package controllers

import (
	initilalizers "main_module/initializers"
	"main_module/models"

	"github.com/gin-gonic/gin"
)

func AddTask(c *gin.Context) {
	// Get data from req body
	var body struct {
		Name      string
		Completed bool
	}
	c.Bind(&body)

	// Create a task
	task := models.Task{
		Name:      body.Name,
		Completed: body.Completed,
	}

	result := initilalizers.DB.Create(&task)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": "Something went wrong",
		})
	}

	c.JSON(200, gin.H{
		"tasks": task,
	})

}

func GetAllTasks(c *gin.Context) {
	// Get all tasks
	var tasks []models.Task
	result := initilalizers.DB.Find(&tasks)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": "Something went wrong",
		})
	}

	// Respond with json
	c.JSON(200, gin.H{
		"tasks": tasks,
	})
}

func GetSingleTask(c *gin.Context) {
	// Get id from request
	id := c.Param("id")

	// Get task
	var task []models.Task
	result := initilalizers.DB.Find(&task, id)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": "Something went wrong",
		})
	}

	// Respond with json
	c.JSON(200, gin.H{
		"tasks": task,
	})
}

func UpdateTask(c *gin.Context) {
	// Get id from request
	id := c.Param("id")

	// Get the data off request body
	var body struct {
		Name      string
		Completed bool
	}

	c.Bind(&body)

	// Find the tasks were updating
	var task []models.Task

	result := initilalizers.DB.Find(&task, id)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": "Something went wrong",
		})
	}

	// Update it
	initilalizers.DB.Model(&task).Updates(map[string]interface{}{"Name": body.Name, "Completed": body.Completed})

	// Respond with json
	c.JSON(200, gin.H{
		"task": task,
	})
}

func DeleteTask(c *gin.Context) {
	// Get the id off the url
	id := c.Param("id")

	// Delete the post
	result := initilalizers.DB.Delete(&models.Task{}, id)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": "Something went wrong",
		})
	}
	// Respond
	c.JSON(200, gin.H{
		"Message": "Task was successfully deleted",
	})
}
