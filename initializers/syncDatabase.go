package initializers

import (
	"golangprogram/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.Task{})
	DB.AutoMigrate(&models.Worker{})
}
