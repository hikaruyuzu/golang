package golangweb

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

// serve file mirip dengan file server akan tetapi dia bisa mengload file sesuai dengan yang kita inginkan(file tertentu)
func ServeFile(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name != "" {
		http.ServeFile(writer, request, "./resources/serveok.html")
	} else {
		http.ServeFile(writer, request, "./resources/servevailed.html")
	}
}

func TestServeFile(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/anime", ServeFile)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// bagamana jika menggunakan embed

//go:embed resources/serveok.html
var success string

//go:embed resources/servevailed.html
var vailed string

func ServeFileEmbed(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name != "" {
		fmt.Fprintf(writer, success)
	} else {
		fmt.Fprint(writer, vailed)
	}
}

func TestServeFileEmbed(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/anime", ServeFileEmbed)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
