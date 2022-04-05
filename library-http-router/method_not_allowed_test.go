package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// digunakan jika method misal kita set di router dengan method get akan tetapi user mengirim dengan method post maka dia
// secara otomatis akan mengembalikan method not allowed akan tetapi jika kamu ingi melakukan costum pada errornya
// kamu bisa menggunakan method not allowed ini

func TestMethodNotAllowed(t *testing.T) {
	router := httprouter.New()

	// handle method not allowed
	router.MethodNotAllowed = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Opppss gabolehh")
	})

	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "nyawww desuuu")
	})
	// NAH DISINI SIMULASINYA
	request := httptest.NewRequest("POST", "http://localhost:8080/", nil)
	record := httptest.NewRecorder()

	router.ServeHTTP(record, request)
	
	response := record.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, "Opppss gabolehh", string(body))
}