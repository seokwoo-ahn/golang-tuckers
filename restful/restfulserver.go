package main

import (
	"golang-tuckers/restful/libs"
	"net/http"
)

func main() {
	http.ListenAndServe(":3000", libs.WebHandler())
}
