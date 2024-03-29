package controllers

import (
	"golangprogram/initializers"
	"golangprogram/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	var updateBody struct {
		Title       string
		Description string
		Status      string
	}

	if err := c.Bind(&updateBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	if len(updateBody.Title) < 3 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title cannot be less than 3 characters"})
		return
	}

	task := models.Task{
		Title:       updateBody.Title,
		Description: updateBody.Description,
		Status:      updateBody.Status,
	}

	result := initializers.DB.Create(&task)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Task could not be created"})
		return
	}

	c.JSON(200, gin.H{
		"task": task,
	})
}

func ListTasks(c *gin.Context) {
	var tasks []models.Task
	initializers.DB.Find(&tasks)

	c.JSON(200, gin.H{
		"tasks": tasks,
	})
}

func GetTask(c *gin.Context) {
	id := c.Param("id")

	var task models.Task
	result := initializers.DB.Where("id = ?", id).First(&task)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(200, gin.H{
		"task": task,
	})
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")

	var updateBody struct {
		Title       string
		Description string
		Status      string
	}
	c.Bind(&updateBody)

	var task models.Task
	result := initializers.DB.First(&task, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Task not found"})
		return
	}

	result = initializers.DB.Model(&task).Updates(models.Task{
		Title:       updateBody.Title,
		Description: updateBody.Description,
		Status:      updateBody.Status,
	})

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Task could not be updated"})
		return
	}

	c.JSON(200, gin.H{
		"newTask": task,
	})
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	initializers.DB.Delete(&models.Task{}, id)
	c.Status(200)
}
