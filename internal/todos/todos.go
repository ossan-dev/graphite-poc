package todos

import (
	"encoding/json"
	"net"
	"net/http"
	"strconv"

	"github.com/ossan-dev/graphitepoc/internal/metrics"
)

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

type TodoHandler struct {
	GraphiteConn net.Conn
}

func NewTodoHandler(conn net.Conn) *TodoHandler {
	return &TodoHandler{
		GraphiteConn: conn,
	}
}

func (t *TodoHandler) GetTodoByID(w http.ResponseWriter, r *http.Request) {
	rawID := r.URL.Query().Get("id")
	if rawID == "" {
		metrics.WriteMetricWithPlaintext(t.GraphiteConn, "webserver.get_todo_by_id.errors.invalid_request", 1.0)
		w.Write([]byte("please provide a TODO ID"))
		return
	}
	id, err := strconv.Atoi(rawID)
	if err != nil {
		metrics.WriteMetricWithPlaintext(t.GraphiteConn, "webserver.get_todo_by_id.errors.invalid_request", 1.0)
		w.Write([]byte("please provide a numeric TODO ID"))
		return
	}
	for _, v := range todos {
		if v.ID == id {
			data, err := json.MarshalIndent(v, "", "\t")
			if err != nil {
				metrics.WriteMetricWithPlaintext(t.GraphiteConn, "webserver.get_todo_by_id.errors.invalid_format", 1.0)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			metrics.WriteMetricWithPlaintext(t.GraphiteConn, "webserver.get_todo_by_id.success", 1.0)
			w.Write(data)
			return
		}
	}
	metrics.WriteMetricWithPlaintext(t.GraphiteConn, "webserver.get_todo_by_id.errors.not_found", 1.0)
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("todo not found"))
}

func (t *TodoHandler) GetTodos(w http.ResponseWriter, r *http.Request) {
	data, err := json.MarshalIndent(todos, "", "\t")
	if err != nil {
		metrics.WriteMetricWithPlaintext(t.GraphiteConn, "webserver.get_todos.errors.invalid_format", 1.0)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	metrics.WriteMetricWithPlaintext(t.GraphiteConn, "webserver.get_todos.success", 1.0)
	w.Write(data)
}
