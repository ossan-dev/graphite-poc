package main

import (
	"net"
	"net/http"

	"github.com/ossan-dev/graphitepoc/internal/config"
	"github.com/ossan-dev/graphitepoc/internal/todos"
)

var todoHandler *todos.TodoHandler

func init() {
	graphiteHost := config.GetEnvOrDefault("GRAPHITE_HOSTNAME", "graphite")
	graphitePort := config.GetEnvOrDefault("GRAPHITE_PORT", "2003")
	conn, err := net.Dial("tcp", net.JoinHostPort(graphiteHost, graphitePort))
	if err != nil {
		panic(err)
	}
	todoHandler = todos.NewTodoHandler(conn)
	if todoHandler == nil {
		panic("could not start the application")
	}
}

func main() {
	defer todoHandler.GraphiteConn.Close()
	http.HandleFunc("/todo", todoHandler.GetTodoByID)
	http.HandleFunc("/todos", todoHandler.GetTodos)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
