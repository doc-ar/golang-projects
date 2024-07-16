package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type todo struct {
	id    int
	value string
}

var todos = []todo{
	{
		id:    0,
		value: "Default Todo Item",
	},
}

func fixTodoIds() {
	for i := range len(todos) {
		todos[i].id = i
	}
}

func addTodo(value string) {
	newTodo := todo{id: len(todos), value: value}
	todos = append(todos, newTodo)
}

func deleteTodo(id int) {
	// Slice deletion function taken from the internet
	todos = append(todos[:id], todos[id+1:]...)
	fixTodoIds()
}

func main() {
	addr := "127.0.0.1:8000"

	if len(os.Args) == 2 {
		addr = os.Args[1]
	}

	fileServer := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	http.HandleFunc("/", renderTodos)
	http.HandleFunc("/newtodo", newTodoHandler)
	http.HandleFunc("/deletetodo", deleteTodoHandler)

	fmt.Printf("Starting server at %s\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
