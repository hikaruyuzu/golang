package golangweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// ini seperti operator if, perbandingan, looping, dan juga nested dari map atau struct
func TemplateAction(writer http.ResponseWriter, request *http.Request) {
	t, err := template.ParseFiles("./templates/if.html")
	if err != nil {
		panic(err)
	}

	t.ExecuteTemplate(writer, "if.html", map[string]interface{}{
		"Title": "Template ACtion",
		"Name":  "Kaguya Shinomiya",
	})
}

// test
func TestTemplateAction(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost/wakaranai", nil)
	record := httptest.NewRecorder()

	TemplateAction(record, request)

	response := record.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func TemplateActionComparation(writer http.ResponseWriter, request *http.Request) {
	t, err := template.ParseFiles("./templates/comparation.html")
	if err != nil {
		panic(err)
	}

	t.ExecuteTemplate(writer, "comparation.html", map[string]interface{}{
		"Title": "Template Action",
		"Value": 90,
	})
}

// test
func TestTemplateActionComparation(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost/wakaranai", nil)
	record := httptest.NewRecorder()

	TemplateActionComparation(record, request)

	response := record.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func TemplateActionRange(writer http.ResponseWriter, request *http.Request) {
	t, err := template.ParseFiles("./templates/range.html")
	if err != nil {
		panic(err)
	}

	t.ExecuteTemplate(writer, "range.html", map[string]interface{}{
		"Title": "Template Action",
		"Hobbies": []string{
			"Makan", "NOnton Anime", "Code",
		},
	})
}

// test
func TestTemplateActionRange(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost/wakaranai", nil)
	record := httptest.NewRecorder()

	TemplateActionRange(record, request)

	response := record.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func TemplateActionWith(writer http.ResponseWriter, request *http.Request) {
	t, err := template.ParseFiles("./templates/with.html")
	if err != nil {
		panic(err)
	}

	t.ExecuteTemplate(writer, "with.html", map[string]interface{}{
		"Title": "Template Action",
		"Address": map[string]string{
			"City":    "Tokyo",
			"Country": "Japan",
		},
	})
}

// test
func TestTemplateActionWith(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost/wakaranai", nil)
	record := httptest.NewRecorder()

	TemplateActionWith(record, request)

	response := record.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
