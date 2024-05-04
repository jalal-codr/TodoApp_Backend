package main

import (
	"errors"
	"log"
	"net/http"

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

func main() {
	// Database conection

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	// uri := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", os.Getenv("host"), os.Getenv("port"), os.Getenv("user"), os.Getenv("password"), os.Getenv("dbname"))
	// db, err := sql.Open("postgres", uri)

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()

	// err = db.Ping()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println("Database connected")

	// if uri == "" {
	// 	log.Fatal("You must set your 'MONGODB_URI' environment variable")

	// }
	// ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	// defer cancel()
	// client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// collection := client.Database("TodoApp").Collection("todos")
	// log.Println(collection)

	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/todos", getTodos)
	router.GET("/todo/:email", getTodo)
	router.POST("/todos", addTodo)

	router.Run("localhost:9090")
}
