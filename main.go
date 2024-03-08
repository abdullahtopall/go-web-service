package main

import (
	"golangprogram/controllers"
	"golangprogram/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	r := gin.Default()
	r.POST("/task", controllers.CreateTask)
	r.GET("/tasks", controllers.ListTasks)
	r.GET("/task/:id", controllers.GetTask)
	r.PUT("/task/:id", controllers.UpdateTask)
	r.DELETE("/task/:id", controllers.DeleteTask)

	r.Run() // listen and serve on 0.0.0.0:8080
}
