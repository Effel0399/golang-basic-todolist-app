package main

import (
	"fmt"
	"net/http"
)

var task1 = "Watch Go crash course"
var task2 = "Take out trash"
var task3 = "Get a donut"
var taskItems = []string{task1, task2, task3}

func main() {
	fmt.Println("##### Welcome to our Todolist App #####")

	// web handlers
	http.HandleFunc("/", helloUser)
	http.HandleFunc("/show-tasks", showTask)

	http.ListenAndServe(":8080", nil) // set address to local and nil for ignore if something happens
}

// getting a request from user and sending response to user
func helloUser(writer http.ResponseWriter, request *http.Request) {
	var greeting = "##### Welcome to our Todolist App #####"
	fmt.Fprintln(writer, greeting)
}

func showTask(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "List of Todos:")
	for index, task := range taskItems {
		fmt.Fprintf(writer, "%d. %s\n", index+1, task)
	}
}

// for adding new task and returning it to the main variable
func addTask(taskItems []string, newTask string) []string {
	var updateTaskItems = append(taskItems, newTask)
	return updateTaskItems
}
