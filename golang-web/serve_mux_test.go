package golangweb

import (
	"fmt"
	"net/http"
	"testing"
)

// serve mux digunakan untuk menangani multiple handler
func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello from home")
	})

	mux.HandleFunc("/anime", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello from anime")
	})

	mux.HandleFunc("/images/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello from images")
	})

	mux.HandleFunc("/images/anilist/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello from Anilist")
	})

	// server
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
