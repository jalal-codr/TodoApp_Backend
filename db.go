package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func databaseConnection() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
	host := os.Getenv("host")
	port, err := strconv.Atoi(os.Getenv("port"))
	if err != nil {
		log.Fatal(err)
	}
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

	//cheack for table

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS Todos (
		id SERIAL PRIMARY KEY,
		email VARCHAR(100),
		name VARCHAR(100) NOT NULL,
		date VARCHAR(100),
		description VARCHAR(200)
		)`)

	if err != nil {
		log.Fatal(err)
	}
	log.Println("Todos table created")

	DB = db

}

func getTodosTable() (*sql.Rows, error) {

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
	host := os.Getenv("host")
	port, err := strconv.Atoi(os.Getenv("port"))
	if err != nil {
		log.Fatal(err)
	}
	user := os.Getenv("user")
	dbname := os.Getenv("dbname")
	password := os.Getenv("password")

	uri := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", uri)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	todo, err := db.Query(`SELECT * FROM Todos`)
	if err != nil {
		return nil, err
	}
	return todo, nil
}
