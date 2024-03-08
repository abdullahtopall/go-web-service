package controllers

import (
	"golangprogram/initializers"
	"golangprogram/models"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	// get data off req body
	var ekle struct {
		Title       string
		Description string
		Status      string
	}
	c.Bind(&ekle)

	// create a post
	task := models.Task{
		Title:       ekle.Title,
		Description: ekle.Description,
		Status:      ekle.Status,
	}

	result := initializers.DB.Create(&task)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// return it
	c.JSON(200, gin.H{
		"task": task,
	})
}

func ListTasks(c *gin.Context) {
	//Get the posts
	var tasks []models.Task
	initializers.DB.Find(&tasks)

	//Respons with them
	c.JSON(200, gin.H{
		"tasks": tasks,
	})
}

func GetTask(c *gin.Context) {
	// get it off url
	id := c.Param("id")

	var task models.Task
	initializers.DB.First(&task, id)

	c.JSON(200, gin.H{
		"task": task,
	})
}

func UpdateTask(c *gin.Context) {
	// get the id off the url
	id := c.Param("id")

	// get the data off req body
	var ekle struct {
		Title       string
		Description string
		Status      string
	}
	c.Bind(&ekle)

	//find the post were updating
	var task models.Task
	initializers.DB.First(&task, id)

	//update it
	initializers.DB.Model(&task).Updates(models.Task{
		Title:       ekle.Title,
		Description: ekle.Description,
		Status:      ekle.Status,
	})

	//respond with it
	c.JSON(200, gin.H{
		"newTask": task,
	})
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	initializers.DB.Delete(&models.Task{}, id)
	c.Status(200)
}
