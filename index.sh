#! /bin/bash

# go run main.go db.go createTodo.go getTodos.go


# echo "$go_files"

# ls
# echo "$files"
go_files=$(ls *.go)
go run $go_files