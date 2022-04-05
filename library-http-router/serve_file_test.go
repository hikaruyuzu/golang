package main

import (
	"embed"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"
)

// sama seperti serve file di mux akan tetapi lebh simple
//go:embed resources
var resources embed.FS

func TestSErveFile(t *testing.T) {
	// akses usb directory
	dictionary, _ := fs.Sub(resources, "resources")

	router := httprouter.New()
	router.ServeFiles("/files/*filepath", http.FS(dictionary))

	request := httptest.NewRequest("GET", "http://localhost:3000/files/hello.html", nil)
	record := httptest.NewRecorder()

	router.ServeHTTP(record, request)

	response := record.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))

}
