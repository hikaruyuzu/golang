package golangweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// untuk passing data yang lebih kompleks
func TemplateDataMap(writer http.ResponseWriter, request *http.Request) {
	t, err := template.ParseFiles("./templates/data.html")
	if err != nil {
		panic(err)
	}

	user := map[string]interface{}{
		"Title": "Template Data Map",
		"Name":  "Kaguya shinomiya",
		"Address": map[string]string{
			"City": "Tokyo",
		},
	}

	t.ExecuteTemplate(writer, "data.html", user)
}

func TestTemplateData(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/kawai?name=Kaguya-chan", nil)
	record := httptest.NewRecorder()

	TemplateDataMap(record, request)

	response := record.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

// bisa juga mengguanakn struct
type Address struct {
	City string
}

type user struct {
	Title   string
	Name    string
	Address Address
}

func TemplateDataStruct(writer http.ResponseWriter, request *http.Request) {
	t, err := template.ParseFiles("./templates/data.html")
	if err != nil {
		panic(err)
	}

	user := user{
		Title: "Template data struct",
		Name:  "Ranko kanzaki",
		Address: Address{
			City: "Hokkaido",
		},
	}
	t.ExecuteTemplate(writer, "data.html", user)
}

func TestTemplateDataStruct(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/kawai?name=Kaguya-chan", nil)
	record := httptest.NewRecorder()

	TemplateDataStruct(record, request)

	response := record.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
