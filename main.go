package main

import (
	"golangprogram/controllers"
	"golangprogram/initializers"
	"golangprogram/models"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {

	workerCount := 5 // İstediğiniz kadar iş parçacığı sayısı
	workerPool := models.WorkerPool{
		Workers:    make([]*models.Worker, workerCount),
		TaskQueue:  make(chan models.Task, 100), // Puffer boyutunu ihtiyaca göre ayarlayın
		QuitSignal: make(chan bool),
	}

	for i := 0; i < workerCount; i++ {
		workerPool.Workers[i] = &models.Worker{TaskCh: make(chan models.Task)}
	}

	workerPool.Start()

	r := gin.Default()
	r.POST("/task", controllers.CreateTask)
	r.GET("/tasks", controllers.ListTasks)
	r.GET("/task/:id", controllers.GetTask)
	r.PUT("/task/:id", controllers.UpdateTask)
	r.DELETE("/task/:id", controllers.DeleteTask)

	r.Run() // listen and serve on 0.0.0.0:8080
}
