package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// membuat httphandle
// router adalah kontrak turunan dari http handler jadi yang bisa dilakukan di handler bisa kita lakukan juga di router
// nah di sini kita bisa memiliki parameter tambahan di handle berupa params
// dan sebelu melakukan router kita bisa menentukan akan menggunakan method apa bak get post put option delete etc

func main() {
	router := httprouter.New()
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprintf(writer, "Hello from kaguya shinomiya")
	})
	
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
