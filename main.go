package main

import (
	"fmt"
)

func main() {
	var task1 = "Watch Go crash course"
	var task2 = "Take out trash"
	var task3 = "Buy a donut"
	// maxTask := 10

	var taskItems = [3]string{task1, task2, task3}

	fmt.Println("##### Welcome to our Todolist App #####")

	i := 0
	taskCount := 1
	for i <= 2 {
		fmt.Print(taskCount)
		fmt.Println(".", taskItems[i])
		i++
		taskCount++
	}

	for _, task := range taskItems {
		fmt.Println(">", task)
	}
}
