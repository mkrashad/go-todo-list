package controllers

import (
	initilalizers "main_module/initializers"
	"main_module/models"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	// Get data from req body
	var body struct {
		FirstName string
		LastName  string
		Email     string
	}
	c.Bind(&body)

	// Create a user
	user := models.User{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
	}

	result := initilalizers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": "Something went wrong",
		})
	}

	c.JSON(200, gin.H{
		"users": user,
	})

}

func GetAllUsers(c *gin.Context) {
	// Get all users
	var users []models.User
	result := initilalizers.DB.Find(&users)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": "Something went wrong",
		})
	}

	// Respond with json
	c.JSON(200, gin.H{
		"users": users,
	})
}

func GetSingleUser(c *gin.Context) {
	// Get id from request
	id := c.Param("id")

	// Get user
	var user []models.User
	result := initilalizers.DB.Find(&user, id)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": "Something went wrong",
		})
	}

	// Respond with json
	c.JSON(200, gin.H{
		"users": user,
	})
}

func UpdateUser(c *gin.Context) {
	// Get id from request
	id := c.Param("id")

	// Get the data off request body
	var body struct {
		FirstName string
		LastName  string
		Email     string
	}

	c.Bind(&body)

	// Find the users were updating
	var user []models.User

	result := initilalizers.DB.Find(&user, id)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": "Something went wrong",
		})
	}

	// Update it
	initilalizers.DB.Model(&user).Updates(map[string]interface{}{"FirstName": body.FirstName, "LastName": body.LastName, "Email": body.Email})

	// Respond with json
	c.JSON(200, gin.H{
		"users": user,
	})
}

func DeleteUser(c *gin.Context) {
	// Get the id off the url
	id := c.Param("id")

	// Delete the user
	result := initilalizers.DB.Delete(&models.User{}, id)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": "Something went wrong",
		})
	}
	// Respond
	c.JSON(200, gin.H{
		"Message": "User was successfully deleted",
	})
}
