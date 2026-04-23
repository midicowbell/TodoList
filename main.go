package main

import (
	"fmt"
	"study/http"
	"study/todo"
)

func main() {
	todolist := todo.NewList()
	HTTPHandlers := http.NewHTTPHandlers(todolist)
	server := http.NewHTTPServer(HTTPHandlers)
	if err := server.StartServer(); err != nil {
		fmt.Println("failed to start HTTP server: ", err)
	}
}
