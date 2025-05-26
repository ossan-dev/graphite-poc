package todos

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"time"
)

var GraphiteConn net.Conn

type todo struct {
	ID          int
	Name        string
	IsCompleted bool
}

var todos = [3]todo{
	{ID: 1, Name: "production service", IsCompleted: false},
	{ID: 2, Name: "test service", IsCompleted: false},
	{ID: 3, Name: "graphite docker file", IsCompleted: false},
}

func writeMetric(name string, value float64) {
	if _, err := GraphiteConn.Write([]byte(fmt.Sprintf("%s %f %d\n", name, value, time.Now().Unix()))); err != nil {
		fmt.Println("error while wrapping metrics to Graphite:", err.Error())
	}
}

func GetTodoByID(w http.ResponseWriter, r *http.Request) {
	rawID := r.URL.Query().Get("id")
	if rawID == "" {
		writeMetric("webserver.get_todo_by_id.errors.invalid_request", 1.0)
		w.Write([]byte("please provide a TODO ID"))
		return
	}
	id, err := strconv.Atoi(rawID)
	if err != nil {
		writeMetric("webserver.get_todo_by_id.errors.invalid_request", 1.0)
		w.Write([]byte("please provide a numeric TODO ID"))
		return
	}
	for _, v := range todos {
		if v.ID == id {
			data, err := json.MarshalIndent(v, "", "\t")
			if err != nil {
				writeMetric("webserver.get_todo_by_id.errors.invalid_format", 1.0)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			writeMetric("webserver.get_todo_by_id.success", 1.0)
			w.Write(data)
			return
		}
	}
	writeMetric("webserver.get_todo_by_id.errors.not_found", 1.0)
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("todo not found"))
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	data, err := json.MarshalIndent(todos, "", "\t")
	if err != nil {
		writeMetric("webserver.get_todos.errors.invalid_format", 1.0)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	writeMetric("webserver.get_todos.success", 1.0)
	w.Write(data)
}
