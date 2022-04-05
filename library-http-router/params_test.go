package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"testing"
)

// params digunakan untuk menetapkan data routing yang dinamis agar bisa berubaha ubah
// misal https://product/:id, nah jika kita ingin data id bisa berubaha ubah maka kita bisa menggunakan fitur dari params ini

func ParamsRouter(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// disini kita bisa ambil nilai dari path yang dibuat dinamis yaitu :idproduct
	id := params.ByName("idproduct")
	if id != "" {
		fmt.Fprintf(writer, "Product id %s", id)
	} else {
		fmt.Fprint(writer, "Konichiwa")
	}

}

func TestRouterParams(t *testing.T) {
	router := httprouter.New()
	// nah artinya disini :idproduct dapat berubah ubah valuenya(dinamis), dan nilainya akan di tampung di params
	router.GET("/product/:idproduct", ParamsRouter)

	//request := httptest.NewRequest("GET", "http://localhost/3000/product/:001", nil)
	//record := httptest.NewRecorder()

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
