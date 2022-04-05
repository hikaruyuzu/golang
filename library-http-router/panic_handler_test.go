package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"testing"
	"text/template"
)

// untuk membuat handler jika terjadi panic di router tanpa menggunakna middleware

func TestPanicHandler(t *testing.T) {

	router := httprouter.New()
	// handler panic
	router.PanicHandler = func(writer http.ResponseWriter, request *http.Request, errorMessage interface{}) {
		handlePanic := map[string]interface{}{
			"error": errorMessage,
		}

		t, err := template.ParseFiles("./resources/handle.html")
		if err != nil {
			panic(err)
		}
		t.Execute(writer, handlePanic)

	}

	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		panic("Oopss something hapenned keep calm ok!")
	})

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	server.ListenAndServe()
}
