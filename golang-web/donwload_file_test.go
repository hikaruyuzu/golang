package golangweb

import (
	"fmt"
	"net/http"
	"testing"
)

// kamu bisa dwnload file dan render secara langsung dengan menambahkan http header berupa "Content-Disposition" ; "attachment ; filename"
func DownloadFile(writer http.ResponseWriter, request *http.Request) {
	fileName := request.URL.Query().Get("file")

	if fileName == "" {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(writer, "STATUS BAD REQUEST")
		return
	}

	// untuk membuat auto download tanpa render
	writer.Header().Add("Content-Disposition", "attachment; filename=\""+fileName+"\"")
	// render file download
	http.ServeFile(writer, request, "./form_upload/"+fileName)
}

func TestDownloadFile(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", DownloadFile)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
