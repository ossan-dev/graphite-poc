package main

import (
	"net/http"

	"github.com/ossan-dev/graphitepoc/internal/todos"
)

func main() {
	http.HandleFunc("/todo", todos.GetTodoByID)
	http.HandleFunc("/todos", todos.GetTodos)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
