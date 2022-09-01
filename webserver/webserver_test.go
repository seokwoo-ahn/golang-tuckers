package main

import (
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
