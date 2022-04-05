package golangweb

import (
	"fmt"
	"net/http"
	"testing"
)

// midleware adalah suatu konsep dimana kita bisa melakukan sesuatu sebelum dan sesudah eksekusi handler
// misal kita ingin mengecek apakah user sudah login atau belum, nah ini akan capek sekali kalua kita cek di setiap handler
// maka dari itu kita bisa manfaatkan middleware disini
// middleware juga bisa digunakan sebagai error handler agar aplikasi tidak mati jika terjadi error(panic)
// di golang tidak ada perintah khusus untuk menangani middleware jadi kita harus buat manual

type LogMiddleware struct {
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// lakukan sesuatu sebelum eksekusi handler
	fmt.Println("Before execute middleware")
	middleware.Handler.ServeHTTP(writer, request)
	// lakuka sesuatu sesudah melakukan handler
	fmt.Println("After execute middleware")
}

// disini kita juga bisa membuat handle untuk error(error handler)
// misal terjadi error di code program kita yang mengembalikan panic, maka kita bisa rcoevr panic itu di dalam middleware
// disarankan untuk membuat handler middleware untuk eror setiap membuat handler
type ErrorHandler struct {
	Handler http.Handler
}

func (errorHandler *ErrorHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// ini adalah handler untuk panic
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Error Handler Execute")
			writer.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(writer, "Error :%s ", err)
		}
	}()

	errorHandler.Handler.ServeHTTP(writer, request)
	// after error handler
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	// disini sy menggunakan anonnymus function agar lebih simple, akan tetapi kamu bisa memasukan handler apapun disini
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Handler execute")
		fmt.Fprintf(writer, "Handler execute")
	})
	mux.HandleFunc("/ara", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Araa kawaii execute")
		fmt.Fprintf(writer, "Ara kawaii execute yaaa")
	})

	mux.HandleFunc("/panic", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Panic desuu")
		// nah panic disini akan di tangkap didalam middleware
		panic("Upps something happened, keep calm :)")
	})

	logMiddleware := &LogMiddleware{
		Handler: mux,
	}
	// ini akan menjadi top middleware untuk handle error, dimana setelah itu dia akan lanjut ke logMiddleware
	errorHandler := &ErrorHandler{
		Handler: logMiddleware,
	}
	// baca pembacaan ekseksui server, ke handler dari sisni
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: errorHandler,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
