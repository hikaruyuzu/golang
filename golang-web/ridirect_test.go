package golangweb

import (
	"fmt"
	"net/http"
	"testing"
)

// digunakan untuk berpindah halaman(ridirect) misal setelah login, etc
// secara default browser sudah memiliki standart untuk ridirect yaitu membuat header path dan mengirimkan return code 3xx
// akan tetapi di golang itu semua sudah otomatis jadi kita hanya perilu memnggil function redirect

func RedirectTo(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello redirect")
}

func RedirectFrom(writer http.ResponseWriter, request *http.Request) {
	// some logic here
	http.Redirect(writer, request, "/redirect-to", http.StatusTemporaryRedirect)
}

func RedirectOut(writer http.ResponseWriter, request *http.Request) {
	// some logic here
	http.Redirect(writer, request, "https://www.facebook.com/profile.php?id=100066661175327", http.StatusTemporaryRedirect)
}

func TestRedirect(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/redirect-from", RedirectFrom)
	mux.HandleFunc("/redirect-to", RedirectTo)
	mux.HandleFunc("/ara-ara", RedirectOut)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
