package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"testing"
	"text/template"
)

// digunakan untuk menghandle costum not found jika router tidak tersedia

func TestHandleNOtFound(t *testing.T) {
	router := httprouter.New()
	// handler panic, jika path tidak tersedia maka dia akan secara otomatis di eksekusi
	router.NotFound = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		handlePanic := map[string]interface{}{
			"error": http.StatusNotFound,
		}

		t, err := template.ParseFiles("./resources/handle.html")
		if err != nil {
			panic(err)
		}
		t.Execute(writer, handlePanic)
	})
	server := http.Server{

		Addr:    "localhost:3000",
		Handler: router,
	}

	server.ListenAndServe()
}
