package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// cara cara penulisan pattern dalam penmbuatan routing di golang
// router pattern by named parameter, catch all router parameter

func TestPatternNamedParameter(t *testing.T) {
	router := httprouter.New()
	router.GET("/product/:id/item/:idItem", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		id := params.ByName("id")
		idItem := params.ByName("idItem")

		fmt.Fprintf(writer, "product with id %s and id item %s", id, idItem)
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/product/1/item/1", nil)
	record := httptest.NewRecorder()

	router.ServeHTTP(record, request)

	response := record.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, "product with id 1 and id item 1", string(body))

}

// kalau catch parameter kita ga peduli di belakangnya ada berapa folder akan tetap bisa di akses
func TestPatternCatchParameter(t *testing.T) {
	router := httprouter.New()
	router.GET("/anime/*movie", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		movie := params.ByName("movie")
		fmt.Fprintf(writer, "anime%s", movie)
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/anime/romance/kaguyasamaloveiswar.mp4", nil)
	record := httptest.NewRecorder()

	router.ServeHTTP(record, request)

	response := record.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, "anime/romance/kaguyasamaloveiswar.mp4", string(body))
	fmt.Println(string(body))
}
