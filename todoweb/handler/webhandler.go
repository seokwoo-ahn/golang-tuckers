package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"sort"
	"strconv"

	todo "golang-tuckers/todoweb/struct"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

var RD *render.Render

func WebHandler() http.Handler {
	todo.TodoMap = make(map[int]todo.Todo)
	mux := mux.NewRouter()
	mux.Handle("/", http.FileServer(http.Dir("public")))
	mux.HandleFunc("/todos", GetTodoListHandler).Methods("GET")
	mux.HandleFunc("/todos", PostTodoHandler).Methods("POST")
	mux.HandleFunc("/todos/{id:[0-9]+}", RemoveTodoHandler).Methods("DELETE")
	mux.HandleFunc("/todos/{id:[0-9]+}", UpdateTodoHandler).Methods("PUT")
	return mux
}

func GetTodoListHandler(w http.ResponseWriter, r *http.Request) {
	list := make(todo.Todos, 0)
	for _, todo := range todo.TodoMap {
		list = append(list, todo)
	}
	sort.Sort(list)
	RD.JSON(w, http.StatusOK, list)
}

func PostTodoHandler(w http.ResponseWriter, r *http.Request) {
	var newTodo todo.Todo
	err := json.NewDecoder(r.Body).Decode(&newTodo)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	todo.LastId++
	todo.TodoMap[todo.LastId] = newTodo
	RD.JSON(w, http.StatusCreated, newTodo)
}

func RemoveTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	if _, ok := todo.TodoMap[id]; ok {
		delete(todo.TodoMap, id)
		RD.JSON(w, http.StatusOK, todo.Success{Success: true})
	} else {
		RD.JSON(w, http.StatusNotFound, todo.Success{Success: false})
	}
}

func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var newTodo todo.Todo
	err := json.NewDecoder(r.Body).Decode(&newTodo)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	if outdated, ok := todo.TodoMap[id]; ok {
		outdated.Name = newTodo.Name
		outdated.Completed = newTodo.Completed
		RD.JSON(w, http.StatusOK, todo.Success{Success: true})
	} else {
		RD.JSON(w, http.StatusBadRequest, todo.Success{Success: false})
	}
}
