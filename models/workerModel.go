package models

import (
	"fmt"
	"sync"

	"gorm.io/gorm"
)

type Worker struct {
	gorm.Model
	TaskCh chan Task
}

func (w *Worker) Start(taskQueue chan Task, wg *sync.WaitGroup, quitSignal chan bool) {
	defer wg.Done()

	for {
		select {
		case task := <-taskQueue:
			// Burada işlemleri gerçekleştirin, örneğin task'ı işleyerek
			// Eğer bir hata olursa, uygun şekilde işleyin
			fmt.Printf("Worker %d processing task %d with params: %v\n", w.ID, task.ID, task.Params)
		case <-quitSignal:
			// Çıkış sinyali alındığında iş parçacığı kapatılır
			return
		}
	}
}
