package utils


// Task represents a task in the todo list.
type Task struct {
	ID int
	Title string
	Status bool
}

type TodoList struct {
	Tasks []Task
	NextID int
}

