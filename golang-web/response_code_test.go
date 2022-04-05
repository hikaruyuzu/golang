package golangweb

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	name := request.URL.Query().Get("name")

	if name != "" {
		writer.WriteHeader(http.StatusOK)
		fmt.Fprintf(writer, "Hello %s", name)
	} else {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, "name is empty")
	}
}

// test
func TestStatusCode(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost/?name=kaguya-chan", nil)
	record := httptest.NewRecorder()

	ResponseCode(record, request)

	response := record.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	status := response.Status
	code := response.StatusCode

	t.Run("success", func(t *testing.T) {
		assert.Equal(t, "name is empty", string(body))
		assert.Equal(t, http.StatusBadRequest, status)
		assert.Equal(t, 400, code)
	})

	t.Run("success", func(t *testing.T) {
		assert.Equal(t, "Hello kaguya-chan", string(body))
		assert.Equal(t, "200 OK", status)
		assert.Equal(t, 200, code)
	})
}
