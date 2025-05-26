package main

import (
	"fmt"
	"net"
	"net/http"

	"github.com/ossan-dev/graphitepoc/internal/todos"
)

func init() {
	var err error
	// HACK: fix this to make it work with IPv6
	todos.GraphiteConn, err = net.Dial("tcp", fmt.Sprintf("%s:%s", "0.0.0.0", "2003"))
	if err != nil {
		panic(err)
	}
}

func main() {
	defer todos.GraphiteConn.Close()
	http.HandleFunc("/todo", todos.GetTodoByID)
	http.HandleFunc("/todos", todos.GetTodos)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
