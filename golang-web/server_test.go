package golangweb

import (
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	// create server
	server := http.Server{
		Addr: "localhost:8080",
	}

	// run server
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
