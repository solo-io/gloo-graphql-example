package main

import (
	"context"
	"fmt"

	todo "github.com/solo-io/gloo-graphql-example/code/todo-app/server"
)

func main() {
	server := todo.NewTodoServer("8080")
	errs, err := server.Start(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("Todo app running on localhost:8080")
	err = <-errs
	panic(err)
}
