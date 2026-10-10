package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type MessageResponce struct {
	Status string `json:"status"`
	Content string `json:"content"`
}

type Task struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
}

func main() {

	serveMux := http.NewServeMux()

	// Simple Endpoint
	serveMux.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		
		msg := MessageResponce{Status: "success", Content: "This is my first http server in go !"}
		json.NewEncoder(w).Encode(msg)
	})

	// Endpoint with path parameter
	serveMux.HandleFunc("GET /hello-from/{user}", func(w http.ResponseWriter, r *http.Request) {
		user := r.PathValue("user")

		w.Header().Set("Content-type", "application/json")
		response := map[string]string{"user":user, "message":"This is my second enpoint in go with path parameter !"}
		json.NewEncoder(w).Encode(response)
	})

	serveMux.HandleFunc("POST /new-task/", func(w http.ResponseWriter, r *http.Request) {

		var newTask Task

		err := json.NewDecoder(r.Body).Decode(&newTask)
		if err != nil{
			w.Header().Set("Content-type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			responce := map[string]string{"message":"Invalid json body"}
			json.NewEncoder(w).Encode(responce)
		}

		fmt.Println("Successfully created new task !")

		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusCreated)

		response := map[string]any{
			"status":  "success",
			"message": "Task created successfully",
			"data":    newTask,
		}
		json.NewEncoder(w).Encode(response)

	})

	// Start the server
	fmt.Println("Server is running at http://localhost:8080")
	if err:= http.ListenAndServe(":8080", serveMux) 

	err != nil{
		log.Fatalf("Server failed to start: %v", err)
	}


}