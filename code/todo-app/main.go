package main

import (
	"context"

	todo "github.com/solo-io/gloo-graphql-example/code/todo-app/server"
)

func main() {
	server := todo.NewTodoServer("8080")
	errs, err := server.Start(context.Background())
	if err != nil {
		panic(err)
	}
	err = <-errs
	panic(err)
}
