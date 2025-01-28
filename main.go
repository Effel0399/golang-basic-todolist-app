package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("##### Welcome to our Todolist App #####")

	http.HandleFunc("/", helloUser)

	http.ListenAndServe(":8080", nil) // set address to local and nil for ignore if something happens
}

// getting a request from user and sending response to user
func helloUser(writer http.ResponseWriter, request *http.Request) {
	var greeting = "Hello user, welcome to our Todolist App!"
	fmt.Fprintln(writer, greeting)
}

// for printing list of tasks
func printTasks(taskItems []string) {
	fmt.Println("List of Todos:")
	// printing arrays with numbering
	for index, task := range taskItems {
		//format with printf
		fmt.Printf("%d. %s\n", index+1, task)
	}
}

// for adding new task and returning it to the main variable
func addTask(taskItems []string, newTask string) []string {
	var updateTaskItems = append(taskItems, newTask)
	return updateTaskItems
}
