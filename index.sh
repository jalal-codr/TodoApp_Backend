#! /bin/bash


go_files=$(find . -maxdepth 2 -type f -name "*.go" -exec basename {} \;)

go run $go_files