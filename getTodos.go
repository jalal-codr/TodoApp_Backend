package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getTodosTable() ([]Todo, error) {

	todo, err := DB.Query("SELECT * FROM Todos")
	if err != nil {
		return nil, err
	}

	var mylist []Todo
	for todo.Next() {

		var value Todo

		err = todo.Scan(&value.Id, &value.Email, &value.Name, &value.Date, &value.Description)
		if err != nil {
			return nil, err
		}
		mylist = append(mylist, value)
	}

	err = todo.Err()
	if err != nil {
		return nil, err
	}

	return mylist, nil
}

func getTodos(context *gin.Context) {

	todo, err := getTodosTable()
	if err != nil {
		log.Fatal(err)
	}

	context.IndentedJSON(http.StatusOK, todo)

}

func getMyTodosTable(email string) ([]Todo, error) {

	statement := ("SELECT * FROM Todos WHERE email =  $1")
	todos, err := DB.Query(statement, &email)
	if err != nil {
		return nil, err
	}
	var mylist []Todo
	for todos.Next() {

		var value Todo

		err = todos.Scan(&value.Id, &value.Email, &value.Name, &value.Date, &value.Description)
		if err != nil {
			return nil, err
		}
		mylist = append(mylist, value)
	}

	err = todos.Err()
	if err != nil {
		return nil, err
	}

	return mylist, nil
}

func getMyTodos(context *gin.Context) {
	email := context.Param("email")

	todos, err := getMyTodosTable(email)
	if err != nil {
		log.Fatal(err)
	}

	context.IndentedJSON(http.StatusOK, todos)
}
