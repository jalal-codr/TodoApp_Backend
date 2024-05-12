package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Todo struct {
	Email       string `json:"email"`
	Id          string `json:"id"`
	Name        string `json:"name"`
	Date        string `json:"date"`
	Description string `json:"description"`
}

func main() {
	//DB connection
	databaseConnection()
	// closeDB()

	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/todos", getTodos)
	router.GET("/myTodos/:email", getMyTodos)
	router.POST("/todos", createTodo)

	router.Run("localhost:9090")
}
