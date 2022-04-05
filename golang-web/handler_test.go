package golangweb

import (
	_ "embed"
	"net/http"
	"testing"
)

// //go:embed files/moetachan3.jpg
var moetachan []byte

// handler digunakan untuk menerimamna request dan mengembalikan response ke client
func TestHandler(t *testing.T) {
	// membuat handler
	var handler http.HandlerFunc = func(writter http.ResponseWriter, request *http.Request) {
		// logic web
		writter.Write(moetachan)
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
