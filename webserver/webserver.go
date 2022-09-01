package main

import (
	"fmt"
	"golang-tuckers/webserver/libs"
	"net/http"
)

func main() {
	http.HandleFunc("/bar", libs.BarHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})
	http.ListenAndServe(":3000", nil)
}
