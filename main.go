package main

import (
	"fmt"
)

func main() {
	var task1 = "Watch Go crash course"
	var task2 = "Take out trash"
	var task3 = "Buy a donut"

	// basic array
	var taskItems = []string{task1, task2, task3}

	fmt.Println("##### Welcome to our Todolist App #####")
	printTasks(taskItems)
	fmt.Println()
	taskItems = addTask(taskItems, "do morning run")
	fmt.Println("Updated list")
	printTasks(taskItems)
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
