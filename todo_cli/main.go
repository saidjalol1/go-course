package main

import (
	"fmt"
	"time"
	"github.com/saidjalol1/todo_cli/utils"
)

func main(){

	myTodoList := &utils.TodoList{
		Tasks: []utils.Task{},
		NextID: 1,
	}

	fmt.Println("Starting Todo CLI Application ....")
	time.Sleep(1000 * time.Millisecond)

	// 2 Test 
	fmt.Println("Adding some tasks to test ....")
	myTodoList.AddTask("Finishing Todo CLI application.")
	myTodoList.AddTask("Building simple REST-API application in GO.")
	myTodoList.AddTask("Building full complex back-end system in GO.")

	// 3 Listing
	myTodoList.ListTasks()
	myTodoList.MarkAsCompleted(2)
	myTodoList.ListTasks()

	myTodoList.DeleteTask(56)
	myTodoList.ListTasks()

	myTodoList.DeleteTask(1)
	myTodoList.ListTasks()
}