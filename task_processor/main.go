package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"github.com/go-course/task_processor/routes"
	"github.com/go-course/task_processor/schemas"
	"github.com/go-course/task_processor/worker"
)

func main() {
	fmt.Println("Starting server on http://localhost:8080 ...")

	var workerWg sync.WaitGroup

	const numWorkers = 2
	for i := 1; i <= numWorkers; i++ {
		workerWg.Add(1)
		go worker.Worker(i, schemas.JOBS_CHANNEL, &workerWg)
	}

	if err := http.ListenAndServe(":8080", routes.Server); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}