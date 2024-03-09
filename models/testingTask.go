package models

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var router *gin.Engine

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	tearDown()
	os.Exit(code)
}

func setup() {
	dsn := "postgres://uzycahxy:xRWVzVFn6bwauFaJHBxjMFvWu6rN2oKx@rain.db.elephantsql.com/uzycahxy"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Test database connection failed")
	}

	db.AutoMigrate(&Task{})

	gin.SetMode(gin.TestMode)

	router = gin.Default()
}

func tearDown() {
	sqlDB, err := db.DB()
	if err != nil {
		panic("Error getting SQL DB from Gorm")
	}
	sqlDB.Close()
}

func TestCreateTask(t *testing.T) {
	task := Task{
		Title:       "Test Task",
		Description: "Testing the task creation",
		Status:      "Pending",
		Params:      map[string]string{"key": "value"},
	}

	db.Create(&task)

	var savedTask Task
	db.First(&savedTask, task.ID)

	assert.Equal(t, task.ID, savedTask.ID)
	assert.Equal(t, task.Title, savedTask.Title)

}

func TestUpdateTask(t *testing.T) {
	initialTask := Task{
		Title:       "Initial Task",
		Description: "Testing task updates",
		Status:      "Pending",
		Params:      map[string]string{"key": "value"},
	}
	db.Create(&initialTask)

	newTitle := "Updated Task Title"
	db.Model(&initialTask).Update("Title", newTitle)

	newDescription := "Updated Description"
	db.Model(&initialTask).Update("Description", newDescription)

	newStatus := "Updated Status"
	db.Model(&initialTask).Update("Status", newStatus)

	var updatedTask Task
	db.First(&updatedTask, initialTask.ID)

	assert.Equal(t, newTitle, updatedTask.Title)
	assert.Equal(t, newDescription, updatedTask.Description)
	assert.Equal(t, newStatus, updatedTask.Status)

}

func TestDeleteTask(t *testing.T) {
	taskToDelete := Task{
		Title:       "Task to Delete",
		Description: "Testing task deletion",
		Status:      "Pending",
		Params:      map[string]string{"key": "value"},
	}
	db.Create(&taskToDelete)

	db.Delete(&taskToDelete)

	var deletedTask Task
	result := db.First(&deletedTask, taskToDelete.ID)
	assert.ErrorIs(t, result.Error, gorm.ErrRecordNotFound)
}

func TestGetTask(t *testing.T) {
	taskToGet := Task{
		Title:       "Task to Get",
		Description: "Testing task retrieval",
		Status:      "Pending",
		Params:      map[string]string{"key": "value"},
	}
	db.Create(&taskToGet)

	var retrievedTask Task
	db.First(&retrievedTask, taskToGet.ID)
	db.First(&retrievedTask, taskToGet.Title)
	db.First(&retrievedTask, taskToGet.Description)
	db.First(&retrievedTask, taskToGet.Status)

	assert.Equal(t, taskToGet.ID, retrievedTask.ID)
	assert.Equal(t, taskToGet.Title, retrievedTask.Title)
	assert.Equal(t, taskToGet.Description, retrievedTask.Description)
	assert.Equal(t, taskToGet.Status, retrievedTask.Status)
}

func TestListTasks(t *testing.T) {
	task1 := Task{
		Title:       "Task 1",
		Description: "Testing task listing",
		Status:      "Pending",
		Params:      map[string]string{"key": "value1"},
	}
	task2 := Task{
		Title:       "Task 2",
		Description: "Testing task listing",
		Status:      "Completed",
		Params:      map[string]string{"key": "value22"},
	}
	db.Create(&task1)
	db.Create(&task2)

	var tasks []Task
	db.Find(&tasks)

	assert.GreaterOrEqual(t, len(tasks), 2)
}
