package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Todo struct {
	User        string `json:"user"`
	Name        string `json:"name"`
	Date        string `json:"date"`
	Description string `json:"description"`
}

var todos = []Todo{}

func getTodos(context *gin.Context) {

	todo, err := getTodosTable()
	if err != nil {
		log.Fatal(err)
	}

	var mylist []Todo
	for todo.Next() {
		var todo Todo
		mylist = append(mylist, todo)
	}

	context.IndentedJSON(http.StatusOK, mylist)

}

func getTodo(context *gin.Context) {
	email := context.Param("email")
	todo, err := getTodoByEmail(email)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"Message": "todo no found"})
	}
	context.IndentedJSON(http.StatusOK, todo)

}

func getTodoByEmail(email string) (*[]Todo, error) {
	var mylist = []Todo{}
	for i, value := range todos {
		if value.User == email {
			mylist = append(mylist, todos[i])
			// return &todos[i], nil
		}
	}
	if len(mylist) > 0 {
		return &mylist, nil
	}
	return nil, errors.New("no todo found")
}

func addTodo(context *gin.Context) {
	var newTodo Todo
	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)

	context.IndentedJSON(http.StatusCreated, newTodo)
}

func main() {
	//DB connection
	databaseConnection()

	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/todos", getTodos)
	router.GET("/todo/:email", getTodo)
	router.POST("/todos", addTodo)

	router.Run("localhost:9090")
}
