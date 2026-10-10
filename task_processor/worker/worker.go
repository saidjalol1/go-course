package worker

import (
	"github.com/go-course/task_processor/schemas"
	"sync"
	"fmt"
	"time"
)




func Worker(WorkerID int, job <-chan schemas.Task, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range job {
		fmt.Printf("Worker %d processing task: %s\n", WorkerID, task.Title)

		// 1 Update the task status to "in_progress"
		schemas.STORE.Mu.Lock()
		task.Status = "in_progress"
		schemas.STORE.Tasks[task.ID] = task
		schemas.STORE.Mu.Unlock()

		time.Sleep(2 * time.Second) // Simulate task processing time

		// 2 Update the task status to "completed"
		schemas.STORE.Mu.Lock()
		task.Status = "completed"
		schemas.STORE.Tasks[task.ID] = task
		schemas.STORE.Mu.Unlock()

		fmt.Printf("Worker %d completed task: %s\n", WorkerID, task.Title)
	}
}