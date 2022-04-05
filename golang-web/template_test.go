package golangweb

import (
	"embed"
	_ "embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// kamu bisa membuat komunikasi file dengan html secara dynamic artinya dengan data dari database/ response dari server
// dengan menggunakan template file

func SimpleTemplate(writer http.ResponseWriter, request *http.Request) {
	templateString := `<html><body>{{.}}</body></html>`
	// sebelum load kita parse dulu templatenya
	t, err := template.New("kawai").Parse(templateString)
	if err != nil {
		panic(err)
	}
	// versi simple parse file tanpa cek error
	// template.Must(template.New("kawai").Parse(templateString))

	data := request.URL.Query().Get("name")
	templateData := fmt.Sprintf("Hello im, %s", data)
	// templateData adalah data yang dikirim kedalam template
	t.ExecuteTemplate(writer, "kawai", templateData)
}

func TestSimpleTemplate(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/kawai?name=Kaguya-chan", nil)
	record := httptest.NewRecorder()

	SimpleTemplate(record, request)

	response := record.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

// nah kamu juga bisa akses template ini dari file
func TemplateSimpleFile(writer http.ResponseWriter, request *http.Request) {
	// parse file
	t, err := template.ParseFiles("./templates/hello.html")
	if err != nil {
		panic(err)
	}
	// bydefault nama dari template akan emngikuti nama filenya
	t.ExecuteTemplate(writer, "hello.html", "Hello, ranko kanzaki")
}

func TestSimpleTemplateFile(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/kawai", nil)
	record := httptest.NewRecorder()

	TemplateSimpleFile(record, request)

	response := record.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

// nah jika ingin load file > 1 kamu bisa menggunakan parse glob agar tidak capek
func TemplateSimpleFileGlob(writer http.ResponseWriter, request *http.Request) {
	// parse file global bisa parsing > 1
	t, err := template.ParseGlob("./templates/*.html")
	if err != nil {
		panic(err)
	}
	// bydefault nama dari template akan emngikuti nama filenya
	t.ExecuteTemplate(writer, "hello.html", "Hello, kitagawa marin")
}

func TestSimpleTemplateFileGlob(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/kawai", nil)
	record := httptest.NewRecorder()

	TemplateSimpleFileGlob(record, request)

	response := record.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

// kamu juga bisa menggunakan fitur embed disini
//go:embed templates/*.html
var templatesEmbed embed.FS

func TemplateFileEmbed(writer http.ResponseWriter, request *http.Request) {
	t, err := template.ParseFS(templatesEmbed, "templates/*.html")
	if err != nil {
		panic(err)
	}

	t.ExecuteTemplate(writer, "hello.html", "Hello ayang tsukasa")
}

func TestSimpleTemplateFileEmbed(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/kawai", nil)
	record := httptest.NewRecorder()

	TemplateFileEmbed(record, request)

	response := record.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
