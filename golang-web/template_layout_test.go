package golangweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// ini digunakan untuk mengabungkan beberapa elemen html yang berbeda seperti header footer dan content
func TemplateLayout(writer http.ResponseWriter, request *http.Request) {
	t, err := template.ParseFiles(""+
		"./templates_layout/header.html",
		"./templates_layout/layout.html",
		"./templates_layout/footer.html")

	if err != nil {
		panic(err)
	}

	anime := map[string]interface{}{
		"Title": "Template layout",
		"Name":  "Kaguya shinomiya",
		"Anime": "Kaguya sama love is war",
	}
	// layout disini berasal dari define yang bisa digunakan untuk membuat alias untuk nama file
	t.ExecuteTemplate(writer, "layout", anime)
}

// test
func TestTemplateLayoutLive(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/anime", TemplateLayout)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestTemplateLayout(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/kawai", nil)
	record := httptest.NewRecorder()

	TemplateLayout(record, request)

	response := record.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
