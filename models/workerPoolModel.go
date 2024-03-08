package models

import (
	"sync"

	"gorm.io/gorm"
)

type WorkerPool struct {
	gorm.Model
	Workers    []*Worker
	TaskQueue  chan Task
	WaitGroup  sync.WaitGroup
	QuitSignal chan bool
}

func (wp *WorkerPool) SubmitTask(task Task) {
	wp.TaskQueue <- task
}

func (wp *WorkerPool) Start() {
	for _, worker := range wp.Workers {
		go worker.Start(wp.TaskQueue, &wp.WaitGroup, wp.QuitSignal)
	}
}
