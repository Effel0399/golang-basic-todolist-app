package main

import (
	"fmt"
)

func main() {
	var task1 = "Watch Go crash course"
	var task2 = "Take out trash"
	var task3 = "Buy a donut"

	// basic array
	var taskItems = [3]string{task1, task2, task3}

	fmt.Println("##### Welcome to our Todolist App #####")

	// printing arrays with numbering
	for index, task := range taskItems {
		//fmt.Println(index+1, ".", task)
		fmt.Printf("%d. %s\n", index+1, task)
	}
}
