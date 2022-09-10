package main

import (
	"encoding/json"
	"golang-tuckers/todoweb/handler"
	todostruct "golang-tuckers/todoweb/struct"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHandler(t *testing.T) {
	assert := assert.New(t)
	mux := handler.WebHandler()
	todoMap := make(todostruct.Todos, 0)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/todos", strings.NewReader(`{"Name":"musthave"}`))
	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusCreated, res.Code)

	res = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/todos", nil)
	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusOK, res.Code)
	err := json.NewDecoder(res.Body).Decode(&todoMap)
	assert.Nil(err)
	assert.Equal(todoMap[0].Name, "musthave")
	assert.Equal(todoMap[0].ID, 0)
	assert.Equal(todoMap[0].Completed, false)
}

func TestUpdateHandler(t *testing.T) {
	assert := assert.New(t)
	mux := handler.WebHandler()
	todoMap := make(todostruct.Todos, 0)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/todos", strings.NewReader(`{"Name":"musthave"}`))
	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusCreated, res.Code)

	res = httptest.NewRecorder()
	req = httptest.NewRequest("PUT", "/todos/0", strings.NewReader(`{"Name":"piano practice"}`))
	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusOK, res.Code)

	res = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/todos", nil)
	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusOK, res.Code)
	err := json.NewDecoder(res.Body).Decode(&todoMap)
	assert.Nil(err)
	assert.Equal(todoMap[0].Name, "piano practice")
	assert.Equal(todoMap[0].ID, 0)
	assert.Equal(todoMap[0].Completed, false)
}
