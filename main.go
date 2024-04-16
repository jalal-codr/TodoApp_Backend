package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	Email string `json:"Email"`
}

var todos = []todo{
	{Email: "jalal@gmail.com"},
	{Email: "jalal@gmail.com"},
	{Email: "jalal@gmail.com"},
	{Email: "jalal@gmail.com"},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context) {
	var newTodo todo
	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)

	context.IndentedJSON(http.StatusCreated, newTodo)
}

func main() {
	router := gin.Default()

	router.GET("/todos", getTodos)
	router.POST("/todos", addTodo)

	router.Run("localhost:9090")
}
