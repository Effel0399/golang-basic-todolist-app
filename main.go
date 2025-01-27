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
}

func printTasks(taskItems []string) {
	fmt.Println("List of Todos:")
	// printing arrays with numbering
	for index, task := range taskItems {
		//format with printf
		fmt.Printf("%d. %s\n", index+1, task)
	}
}
