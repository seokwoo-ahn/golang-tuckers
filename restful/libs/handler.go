package libs

import (
	"encoding/json"
	"net/http"
	"sort"

	"github.com/gorilla/mux"
)

func WebHandler() http.Handler {
	mux := mux.NewRouter()
	mux.HandleFunc("/students", StudentListHandler).Methods("GET")

	StudentsMap = make(map[int]Student)
	StudentsMap[1] = Student{1, "cranberry", 2, 100}
	StudentsMap[2] = Student{2, "ralo", 29, 24000}

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
