package main

import (
	"golangprogram/controllers"
	"golangprogram/initializers"
	"golangprogram/middleware"
	"golangprogram/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
	swaggerFiles "github.com/swaggo/files"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/example/basic/docs"
)

func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
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

	workerCount := 5
	workerPool := models.WorkerPool{
		Workers:    make([]*models.Worker, workerCount),
		TaskQueue:  make(chan models.Task, 100),
		QuitSignal: make(chan bool),
	}

	for i := 0; i < workerCount; i++ {
		workerPool.Workers[i] = &models.Worker{TaskCh: make(chan models.Task)}
	}

	workerPool.Start()

	r.POST("/task", controllers.CreateTask)
	r.GET("/tasks", controllers.ListTasks)
	r.GET("/task/:id", controllers.GetTask)
	r.PUT("/task/:id", controllers.UpdateTask)
	r.DELETE("/task/:id", controllers.DeleteTask)

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	r.Run()
}

func TestHelloworld(t *testing.T) {
	router := gin.Default()

	router.GET("/helloworld", Helloworld)

	req, err := http.NewRequest("GET", "/helloworld", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	context, _ := gin.CreateTestContext(rr)

	context.Request = req

	Helloworld(context)

	assert.Equal(t, http.StatusOK, rr.Code)

	expected := "\"helloworld\"\n"
	assert.Equal(t, expected, rr.Body.String())
}
