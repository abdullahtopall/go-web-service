package main

import (
	"golangprogram/controllers"
	"golangprogram/initializers"
	"golangprogram/models"
	"net/http"

	swaggerFiles "github.com/swaggo/files"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/example/basic/docs"
)

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {

	r := gin.Default()
	docs.SwaggerInfo_swagger.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		eg := v1.Group("/example")
		{
			eg.GET("/helloworld", Helloworld)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// r.Run(":8080")

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

	// r := gin.Default()
	r.POST("/task", controllers.CreateTask)
	r.GET("/tasks", controllers.ListTasks)
	r.GET("/task/:id", controllers.GetTask)
	r.PUT("/task/:id", controllers.UpdateTask)
	r.DELETE("/task/:id", controllers.DeleteTask)

	r.Run() // listen and serve on 0.0.0.0:8080
}
