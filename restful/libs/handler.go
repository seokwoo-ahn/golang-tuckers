package libs

import (
	"encoding/json"
	"net/http"
	"sort"
	"strconv"

	"github.com/gorilla/mux"
)

func WebHandler() http.Handler {
	mux := mux.NewRouter()
	mux.HandleFunc("/students", StudentListHandler).Methods("GET")

	StudentsMap = make(map[int]Student)
	StudentsMap[1] = Student{1, "cranberry", 2, 100}
	StudentsMap[2] = Student{2, "ralo", 29, 2400}

	mux.HandleFunc("/students/{id:[0-9]+}", GetStudentHandler).Methods("GET")

	return mux
}

func StudentListHandler(w http.ResponseWriter, r *http.Request) {
	list := make(Students, 0)
	for _, student := range StudentsMap {
		list = append(list, student)
	}
	sort.Sort(list)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}

func GetStudentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	student, ok := StudentsMap[id]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}
