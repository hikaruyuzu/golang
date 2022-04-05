package golangweb

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// digunakan untuk menangkap request dari user, dan akan bannyak digunakan

func HelloWaifu(writer http.ResponseWriter, request *http.Request) {
	// mendapatkan data request, dengan query param
	name := request.URL.Query().Get("waifu")
	if name != "" {
		fmt.Fprintf(writer, "Hello ayang %s", name)
	} else {
		fmt.Fprint(writer, "Hallo konichiwa")
	}
	// disini kamu bisa menghubungkannya ke controller/service, untuk melakukan sesuatu

}

// testing query param HelloWaifu
func TestQueryParam(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost/wakaranai?waifu=kaguyachan", nil)
	record := httptest.NewRecorder()

	HelloWaifu(record, request)

	response := record.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	t.Run("test hello waifu success",
		func(t *testing.T) {
			assert.Equal(t, "Hello ayang kaguyachan", string(body))
		})

	t.Run("test hello waifu failed",
		func(t *testing.T) {
			assert.Equal(t, "Hello konichiwa", string(body))
		})

	fmt.Println(string(body))
}

// multiple request param
func KonichiwaMultiple(writer http.ResponseWriter, request *http.Request) {
	// jika ada lebih dari 1 request kamu bisa membuat methodnya sebanyak param request yang ada
	anime := request.URL.Query().Get("anime")
	name := request.URL.Query().Get("name")

	if anime != "" && name != "" {
		fmt.Fprintf(writer, "hello %s from anime %s", name, anime)
	} else {
		fmt.Fprint(writer, "Hello konichiwa")
	}
}

// test multiple request param
func TestMultipleRequestParam(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost/wakaranai?name=kaguya-sama&anime=kaguya-sama-love-is-war", nil)
	record := httptest.NewRecorder()

	KonichiwaMultiple(record, request)

	response := record.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	t.Run("test hello waifu success",
		func(t *testing.T) {
			assert.Equal(t, "hello kaguya-sama from anime kaguya-sama-love-is-war", string(body))
		})

	t.Run("test hello waifu failed",
		func(t *testing.T) {
			assert.Equal(t, "Hello konichiwa", string(body))
		})

	fmt.Println(string(body))
}

// nah bagaimana jika menangkapnya dengan satu param yang sama akan tetapi paramnya lebih dari 1?
// by default request.uRL.Query akan mengembalikan slice of string yang ada di dalam map jadi kita bisa mendaptkan nilainya dengan index map
func HelloKonichiwaMultipleInOneParam(writer http.ResponseWriter, request *http.Request) {
	title := request.URL.Query()
	animeTitle := title["anime"]
	// akan saya tampilkan
	for _, title := range animeTitle {
		fmt.Fprintf(writer, "Hello from anime %s \n", title)
	}
}

func TestMultipleInOnePrameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost/wakaranai?anime=wake-up-girls&anime=kaguya-sama-love-is-war&anime=genjutsu-yuusha", nil)
	record := httptest.NewRecorder()

	HelloKonichiwaMultipleInOneParam(record, request)

	response := record.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
