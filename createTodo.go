package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func createTodo(context *gin.Context) {
	var newTodo Todo
	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	_, err := addTodo(newTodo)
	if err != nil {
		context.IndentedJSON(http.StatusCreated, err)
	}

	context.IndentedJSON(http.StatusCreated, "New todo created")
}

func addTodo(todo Todo) (sql.Result, error) {

	statement := (`INSERT INTO Todos(email,name,date,description) VALUES ($1, $2, $3, $4)`)
	newTodo, err := DB.Exec(statement, todo.Email, todo.Name, todo.Date, todo.Description)
	if err != nil {
		return nil, err
	}
	return newTodo, nil
}
