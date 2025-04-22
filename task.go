package main

import (
	"fmt"
	"time"
)

type Task struct {
	id        int
	title     string
	completed bool
	createdAt time.Time
}

func newTask(id int, title string) Task {
	return Task{
		id:        id,
		title:     title,
		completed: false,
		createdAt: time.Now(),
	}
}

var tasks []Task

func AddTask(title string) {
	id := 1
	userTask := newTask(id, title)
}

func main() {
	var taskTitle string
	fmt.Print("Enter a title for task:")
	fmt.Scanln(&taskTitle)
	AddTask(taskTitle)
}
