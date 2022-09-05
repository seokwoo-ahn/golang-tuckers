package main

import (
	"encoding/json"
	"golang-tuckers/restful/libs"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/students", nil)

	mux := libs.WebHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	var list []libs.Student
	err := json.NewDecoder(res.Body).Decode(&list)
	assert.Nil(err)
	assert.Equal(2, len(list))
}

func TestJsonHandler2(t *testing.T) {
	assert := assert.New(t)

	var student libs.Student
	mux := libs.WebHandler()
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/students/1", nil)

	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusOK, res.Code)
	err := json.NewDecoder(res.Body).Decode(&student)
	assert.Nil(err)
	assert.Equal("cranberry", student.Name)

	res = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/students/2", nil)

	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusOK, res.Code)
	err = json.NewDecoder(res.Body).Decode(&student)
	assert.Nil(err)
	assert.Equal("ralo", student.Name)
}