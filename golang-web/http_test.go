package golangweb

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// digunakan untuk melakukan test terhadap handler

func KonichiwaHandler(writter http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writter, "Konichiwa Onichann")
}

// test handler
func TestHttp(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/konichiwa", nil)
	recorder := httptest.NewRecorder()

	KonichiwaHandler(recorder, request)

	result := recorder.Result()
	body, err := io.ReadAll(result.Body)
	if err != nil {
		panic(err)
	}

	bodySring := string(body)
	assert.Equal(t, "Konichiwa Onichann", bodySring)

}
