package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type todo struct {
	User        string `json:"user"`
	Name        string `json:"name"`
	Date        string `json:"date"`
	Description string `json:"description"`
}

var todos = []todo{}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func getTodo(context *gin.Context) {
	email := context.Param("email")
	todo, err := getTodoByEmail(email)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"Message": "todo no found"})
	}
	context.IndentedJSON(http.StatusOK, todo)

}

func getTodoByEmail(email string) (*[]todo, error) {
	var mylist = []todo{}
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
	var newTodo todo
	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)

	context.IndentedJSON(http.StatusCreated, newTodo)
}
func databaseConnection() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
	host := os.Getenv("host")
	port := 5432
	user := os.Getenv("user")
	dbname := os.Getenv("dbname")
	password := os.Getenv("password")

	uri := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", uri)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database connected")
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
