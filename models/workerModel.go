package models

import (
	"fmt"
	"sync"

	"gorm.io/gorm"
)

type Worker struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
	TaskCh   chan Task
}

func (w *Worker) Start(taskQueue chan Task, wg *sync.WaitGroup, quitSignal chan bool) {
	defer wg.Done()

	for {
		select {
		case task := <-taskQueue:
			fmt.Printf("Worker %d processing task %d with params: %v\n", w.ID, task.ID, task.Params)
		case <-quitSignal:
			return
		}
	}
}
