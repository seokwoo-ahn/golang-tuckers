package libs

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Student struct {
	Name  string
	Age   int
	Score int
}

func JsonHandler(w http.ResponseWriter, r *http.Request) {
	var student = Student{"ralo", 28, 2400}
	data, _ := json.Marshal(student)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data))
}
