package utils

import (
	"fmt"
)


func (list *TodoList) AddTask(title string) {

	newTask := Task{
		ID:     list.NextID,
		Title:  title,
		Status: false,
	}

	list.Tasks = append(list.Tasks, newTask)

	list.NextID++

	fmt.Printf("Successfully added task #%d: '%s'\n", newTask.ID, title)
}


func (list *TodoList) ListTasks(){

	if len(list.Tasks) == 0{
		fmt.Println("You have not added tasks yet !!!")
		return
	}

	fmt.Println("===== Your Tasks =====")

	for _ , task := range list.Tasks{
		status := " "
		if task.Status == false{
			status = "X"
		}else{
			status = "/"
		}
		
		fmt.Printf("[%s] ID: %d - %s\n", status, task.ID, task.Title)
	}
	fmt.Println("======     End    =====")

}


func (list *TodoList) MarkAsCompleted(id int){

	if len(list.Tasks) == 0{
		fmt.Println("You have not added tasks yet !!!")
		return
	}

	for i := range list.Tasks{
		if list.Tasks[i].ID == id{
			list.Tasks[i].Status = true
			fmt.Printf("Success: Your task #%d is marked as completed\n", id)
			return
		}
	}
	fmt.Printf("Error: Your task with ID #%d is not found!\n", id)
}


func (list *TodoList) DeleteTask(id int){

	if len(list.Tasks) == 0{
		fmt.Println("You have not added tasks yet !!!")
		return
	}

	for i := range list.Tasks{
		if list.Tasks[i].ID == id{
			
			list.Tasks = append(list.Tasks[:i], list.Tasks[i+1:]...)

			fmt.Printf("Success: Deleted Task with the id #%d\n", id)
			return
		} 
	}
	fmt.Println("Error: Could not found Task with id #%d", id)

}