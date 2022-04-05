package golangweb

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// header adalah infomasi tambahan di browser yang dikirimkan dari server ke client atau client ke server melalui request atau response writer
func RequestHeader(writer http.ResponseWriter, request *http.Request) {
	// get header from request, dari client
	contentHeader := request.Header.Get("content-type")
	fmt.Fprint(writer, contentHeader)
}

// test request header
func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost/kawai", nil)
	record := httptest.NewRecorder()

	// add header
	request.Header.Add("content-type", "application/json")

	RequestHeader(record, request)

	response := record.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, "application/json", string(body))
}

// response header dari server
func ResponseHeader(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("x-powered-by", "yukisora")
	fmt.Fprint(writer, http.StatusOK)
}

// test response header dari server
func TestResponseHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost/kawaii", nil)
	record := httptest.NewRecorder()

	ResponseHeader(record, request)

	response := record.Result()
	body, _ := io.ReadAll(response.Body)

	header := response.Header.Get("x-powered-by")

	assert.Equal(t, "yukisora", header)
	assert.Equal(t, "200", string(body))

}
