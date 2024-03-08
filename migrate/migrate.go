package main

import (
	"golangprogram/initializers"
	"golangprogram/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.Task{})
}
