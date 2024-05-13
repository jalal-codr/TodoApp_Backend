package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func deleteTodo(context *gin.Context) {
	var todo Todo
	err := context.BindJSON(&todo)
	if err != nil {
		return
	}
	response, err := removeTodo(todo)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, err)
	}
	context.IndentedJSON(http.StatusAccepted, response)

}

func removeTodo(todo Todo) (sql.Result, error) {
	statement := "DELETE FROM Todos WHERE Id = $1"
	response, err := DB.Exec(statement, todo.Id)
	if err != nil {
		return nil, err
	}
	return response, nil
}
