package main

import (
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
var task []Task
var nextId int = 1

// gin router
func main() {
	// create router with gin default
	router := gin.Default()

	// home
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the Task Manager web application!",
		})

	})

	// router for Add a new TRask(with json)
	router.POST("/tasks", func(ctx *gin.Context) {
		var newtask Task

		// Bind JSON input to struct
		if err := ctx.BindJSON(&newtask); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		}

		newtask.ID = nextId
		nextId++
		newtask.Done = false

		task = append(task, newtask)
		ctx.JSON(http.StatusCreated, newtask)

	})

	// router for List all tasks (as json)
	router.GET("/tasks", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, task)
	})
	// listening the server
	router.Run(":8080")
}
