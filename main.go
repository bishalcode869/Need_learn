package main

import (
	"encoding/json"
	"fmt"
)

// now marshalling
type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func main() {
	// struct data
	task := Task{ID: 1, Title: "Learn GO JSON", Done: true}

	// converting it
	jsonData, err := json.Marshal(task)
	if err != nil {
		fmt.Println("Error encoding JSON: ", err)
	}
	fmt.Println(string(jsonData))
}

// // unmarshalling

// type Task struct {
// 	Title string `json: "title"`
// }

// // main function
// func main() {
// 	jsonData := []byte(`{"title": "Learn Gin"}`)

// 	var task Task
// 	err := json.Unmarshal(jsonData, &task)
// 	if err != nil {
// 		fmt.Println("Error decoding json: ", err)
// 	}
// 	fmt.Println("Task Title: ", task.Title)
// }
