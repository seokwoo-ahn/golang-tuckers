package main

import (
	"golang-tuckers/todoweb/handler"
	"log"
	"net/http"

	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

func main() {
	handler.RD = render.New()
	m := handler.WebHandler()
	n := negroni.Classic()
	n.UseHandler(m)

	log.Println("Started App")
	err := http.ListenAndServe(":3000", n)
	if err != nil {
		panic(err)
	}
}
