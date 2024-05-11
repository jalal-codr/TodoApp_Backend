package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Todo struct {
	Id          string `json:id`
	Email       string `json:"email"`
	Name        string `json:"name"`
	Date        string `json:"date"`
	Description string `json:"description"`
}

func main() {
	//DB connection
	databaseConnection()

	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/todos", getTodos)
	router.GET("/myTodos/:email", getMyTodos)
	router.POST("/todos", createTodo)

	router.Run("localhost:9090")
}
