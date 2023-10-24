package pool

import (
	"log"
)

// Worker handles all the work.
type Worker struct {
	ID       int
	taskChan chan *Task
	quit     chan bool
}

// NewWorker returns new instance of worker.
func NewWorker(channel chan *Task, id int) *Worker {
	return &Worker{
		ID:       id,
		taskChan: channel,
		quit:     make(chan bool),
	}
}

// Start starts the worker.
func (wr *Worker) Start() {
	log.Printf("Starting worker %d\n", wr.ID)

	for {
		select {
		case task := <-wr.taskChan:
			process(wr.ID, task)
		case <-wr.quit:
			return
		}
	}
}

// Stop quits the worker.
func (wr *Worker) Stop() {
	log.Printf("Closing worker %d\n", wr.ID)

	go func() {
		wr.quit <- true
	}()
}
