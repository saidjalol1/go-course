package routes


import (
	"github.com/go-course/task_processor/schemas"
	"net/http"
	"encoding/json"
	"time"
	"fmt"
	"strconv"
)

var Server = http.NewServeMux()

func init() {
	
Server.HandleFunc("POST /new-task/", func(w http.ResponseWriter, r *http.Request) {
		var requestBody struct {
			Title string `json:"title"`
		}

		err := json.NewDecoder(r.Body).Decode(&requestBody);
		if err != nil || requestBody.Title == "" {
			w.Header().Set("Content-type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			response := map[string]string{"message": "Invalid JSON body"}
			json.NewEncoder(w).Encode(response)
			return 
		}

		// Create a new task
		taskID := fmt.Sprintf("%d", time.Now().UnixNano())
		newTask := schemas.Task{
			ID:        taskID,
			Title:     requestBody.Title,
			Status:    "pending",
			CreatedAt: time.Now(),
		}

		schemas.STORE.Mu.Lock()
		schemas.STORE.Tasks[taskID] = newTask
		schemas.STORE.Mu.Unlock()

		// Send the task to the jobs channel for processing
		schemas.JOBS_CHANNEL <- newTask

		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(map[string]any{"message": "Task created and sent for processing", "task": newTask})

})


Server.HandleFunc("GET /tasks", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		limit := 10  // default limit
		offset := 0 // default offset

		l := query.Get("limit")
		if  l != "" {
			if parsed, err := strconv.Atoi(l); err == nil {
				limit = parsed
			}
		}

		o := query.Get("offset")
		if o != "" {
			if parsed, err := strconv.Atoi(o); err == nil {
				offset = parsed
			}
		}

		// Read tasks safely from the thread-safe store
		schemas.STORE.Mu.RLock()
		
		// Collect all tasks into a slice for pagination
		var allTasks []schemas.Task
		for _, task := range schemas.STORE.Tasks {
			allTasks = append(allTasks, task)
		}
		
		schemas.STORE.Mu.RUnlock()

		// Simple pagination slicing logic
		start := offset
		if start > len(allTasks) {
			start = len(allTasks)
		}

		end := start + limit
		if end > len(allTasks) {
			end = len(allTasks)
		}

		paginatedTasks := allTasks[start:end]

		// Send back the paginated response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{
			"total":  len(allTasks),
			"limit":  limit,
			"offset": offset,
			"tasks":  paginatedTasks,
		})
	})
}