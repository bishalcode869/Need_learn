package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// first types struct
type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

// store task and next id
var tasks []Task
var nextId int = 1

// gin router
func main() {
	// create router with gin default
	router := gin.Default()

	// router HOME route
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the Task Manager web application!",
		})

	})

	// router POST route
	router.POST("/tasks", func(ctx *gin.Context) {
		var newtask Task

		// Bind JSON input to struct
		if err := ctx.BindJSON(&newtask); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		newtask.ID = nextId
		nextId++
		newtask.Done = false

		tasks = append(tasks, newtask)
		ctx.JSON(http.StatusCreated, newtask)

	})

	// router for List all tasks (as json)
	router.GET("/tasks", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, tasks)
	})

	// router PUT route
	router.PUT("/tasks/:id/Done", func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		var id int
		_, err := fmt.Sscanf(idParam, "%d", &id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
			return
		}

		for i := range tasks {
			if tasks[i].ID == id {
				tasks[i].Done = true
				ctx.JSON(http.StatusOK, tasks[i])
				return
			}
		}
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
	})

	// router DELETE route
	router.DELETE("/tasks/:id", func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		var id int
		_, err := fmt.Sscanf(idParam, "%d", &id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
			return
		}

		index := -1
		for i, t := range tasks {
			if t.ID == id {
				index = i
				break
			}
		}

		if index == -1 {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			return
		}

		tasks = append(tasks[:index], tasks[index+1:]...)
		ctx.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
	})
	// listening the server
	router.Run(":8080")
}
