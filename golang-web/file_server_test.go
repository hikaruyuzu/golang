package golangweb

import (
	"embed"
	"io/fs"
	"net/http"
	"testing"
)

// file server digunakan untuk render/ menampilkan file static berupa html css image js dll
func TestFileServer(t *testing.T) {
	// mengambil data dari package resources
	resources := http.Dir("./resources")
	fileServer := http.FileServer(resources)

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// jika menggunakan embed

//go:embed resources
var resources embed.FS

func TestEmbedFileServer(t *testing.T) {
	// agar tidak perlu mengakses ke folder resource lagi
	subResources, _ := fs.Sub(resources, "resources")
	fileServer := http.FileServer(http.FS(subResources))

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
