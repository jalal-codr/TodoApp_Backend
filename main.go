package main

import (
	"errors"
	"net/http"

	"github.com/gin-contrib/cors"
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
		if value.Email == email {
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
	//Database conection
	// if err := godotenv.Load(); err != nil {
	// 	log.Println("No .env file found")
	// }
	// uri := os.Getenv("MONGODB_URI")
	// if uri == "" {
	// 	log.Fatal("You must set your 'MONGODB_URI' environment variable")

	// }
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	// client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://panda:panda@cluster0.sk272un.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer func() {
	// 	if err := client.Disconnect(ctx); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }()

	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/todos", getTodos)
	router.GET("/todo/:email", getTodo)
	router.POST("/todos", addTodo)

	router.Run("localhost:9090")
}
