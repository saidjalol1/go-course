package schemas

import (
	"time"
	"sync"
)


type Task struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Status string `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}


type TaskStore struct {
	Mu sync.RWMutex
	Tasks map[string]Task
}


var STORE = TaskStore{
	Tasks: make(map[string]Task),
}

var JOBS_CHANNEL = make(chan Task, 100) // Buffered channel to hold tasks





