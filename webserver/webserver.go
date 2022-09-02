package main

import (
	"fmt"
	"golang-tuckers/webserver/libs"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/bar", libs.BarHandler)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})
	mux.HandleFunc("/json", libs.JsonHandler)
	http.ListenAndServe(":3000", mux)
}
