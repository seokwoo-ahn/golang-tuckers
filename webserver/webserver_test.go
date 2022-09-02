package main

import (
	"encoding/json"
	"fmt"
	"golang-tuckers/webserver/libs"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar", nil)
	mux := http.NewServeMux()
	mux.HandleFunc("/bar", libs.BarHandler)
	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusOK, res.Code)
	data, _ := io.ReadAll(res.Body)
	fmt.Println(data)
}

func TestJsonHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/json", nil)

	mux := http.NewServeMux()
	mux.HandleFunc("/json", libs.JsonHandler)
	mux.ServeHTTP(res, req)
	student := new(libs.Student)
	err := json.NewDecoder(res.Body).Decode(student)
	assert.Nil(err)
	assert.Equal("ralo", student.Name)
}
