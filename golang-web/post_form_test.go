package golangweb

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// digunakan untuk mengambil post form dari html, yang mana dia akan ditaruh dibody, tidak seperti get
func FormPost(writer http.ResponseWriter, request *http.Request) {
	// parsing form post dari body request, agar dapat dibaca oleh PostForm
	err := request.ParseForm()
	if err != nil {
		panic(err)
	}
	// cara cepat tanpa parse dulu
	// example := request.PostFormValue("key")

	// kemudian baca
	title := request.PostForm.Get("title")
	name := request.PostForm.Get("name")

	fmt.Fprintf(writer, "Hello, my name is %s from anime %s", name, title)
}

// testing sudah mengikuti standart
func TestFormPost(t *testing.T) {
	postBody := strings.NewReader("title=kaguya-sama-love-is-war&name=kaguya-shinomiya")
	request := httptest.NewRequest(http.MethodPost, "http://localhost/anime", postBody)
	// header sudah menjadi standard
	request.Header.Add("content-type", "application/x-www-form-urlencoded")
	record := httptest.NewRecorder()

	FormPost(record, request)

	response := record.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, "Hello, my name is kaguya-shinomiya from anime kaguya-sama-love-is-war", string(body))
}
