package golangweb

import (
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// cookie digunakan untuk menyimpan data di sisi client

func SetCookie(writer http.ResponseWriter, request *http.Request) {
	cookieName := request.URL.Query().Get("name")
	hash := sha256.New()
	_, err := hash.Write([]byte(cookieName))
	if err != nil {
		panic(err)
	}

	// kamu bisa set cookie disini kebih dari satu
	cookie := new(http.Cookie)
	cookie.Name = "x-yukisora"
	cookie.Value = fmt.Sprintf("%x", hash.Sum(nil))
	cookie.Path = "/"
	cookie.Expires = time.Now().AddDate(1, 0, 0)

	http.SetCookie(writer, cookie)

	fmt.Fprint(writer, "success set cookie name")
}

func GetCookie(writer http.ResponseWriter, request *http.Request) {
	cookies, err := request.Cookie("x-yukisora")

	if err != nil {
		fmt.Fprint(writer, "No cookie")
	} else {
		fmt.Fprintf(writer, "Hash code %s", cookies.Value)
	}
}

func TestCookie(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// unit test cookie
func TestSetCookie(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost/?name=kaguya-shinomiya", nil)
	record := httptest.NewRecorder()

	SetCookie(record, request)

	response := record.Result().Cookies()
	for _, cookie := range response {
		fmt.Printf("cookie %s %s", cookie.Name, cookie.Value)
	}

}

func TestGetCookie(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost/", nil)
	record := httptest.NewRecorder()

	cookie := new(http.Cookie)
	cookie.Name = "x-yukisora"
	cookie.Value = "test"
	request.AddCookie(cookie)

	GetCookie(record, request)

	response := record.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
