package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func editTodo(context *gin.Context) {
	var newTodo Todo
	err := context.BindJSON(&newTodo)
	if err != nil {
		return
	}
	response, err := updateTodo(newTodo)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, err)
	}
	context.IndentedJSON(http.StatusAccepted, response)
}

func updateTodo(todo Todo) (sql.Result, error) {
	satements := "UPDATE Todos SET email = $5, name = $1, description = $2, date = $3 WHERE Id = $4"
	response, err := DB.Exec(satements, todo.Name, todo.Description, todo.Date, todo.Id, todo.Email)
	if err != nil {
		return nil, err
	}
	return response, nil

}
