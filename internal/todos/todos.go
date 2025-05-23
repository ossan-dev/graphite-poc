package todos

import (
	"encoding/json"
	"net/http"
	"strconv"
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

func GetTodoByID(w http.ResponseWriter, r *http.Request) {
	rawID := r.URL.Query().Get("id")
	if rawID == "" {
		w.Write([]byte("please provide a TODO ID"))
		return
	}
	id, err := strconv.Atoi(rawID)
	if err != nil {
		w.Write([]byte("please provide a numeric TODO ID"))
		return
	}
	for _, v := range todos {
		if v.ID == id {
			data, err := json.MarshalIndent(v, "", "\t")
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Write(data)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("todo not found"))
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	data, err := json.MarshalIndent(todos, "", "\t")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(data)
}
