package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

type Todo struct {
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name"`
	Completed bool   `json:"completed,omitempty"`
}

var todoMap map[int]Todo
var lastId int = 0

var RD *render.Render

func WebHandler() http.Handler {
	todoMap = make(map[int]Todo)
	mux := mux.NewRouter()
	mux.Handle("/", http.FileServer(http.Dir("public")))
	mux.HandleFunc("/todos", GetTodoListHandler).Methods("GET")
	mux.HandleFunc("/todos", PostTodoHandler).Methods("POST")
	mux.HandleFunc("/todos/{id:[0-9]+}", RemoveTodoHandler).Methods("DELETE")
	mux.HandleFunc("/todos/{id:[0-9]+}", UpdateTodoHandler).Methods("PUT")
	return mux
}

func GetTodoListHandler(w http.ResponseWriter, r *http.Request) {

}

func PostTodoHandler(w http.ResponseWriter, r *http.Request) {

}

func RemoveTodoHandler(w http.ResponseWriter, r *http.Request) {

}

func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {

}
