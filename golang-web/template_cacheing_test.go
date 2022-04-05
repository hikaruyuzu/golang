package golangweb

import (
	"embed"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

// cacheing digunakan agar code program kita semakin cepat
// tidak perlu parse file terus menerus di dalam handler, cukup 1 kali saja
// gunakan cacheing agar code lebih baik dan berjalan cepat, data akan di simoa di memory

//go:embed template_function
var embedTemplateFunc embed.FS

// nah ini adalah pembutaan template caching di awal
var templateFuncCacheing = template.Must(template.ParseFS(embedTemplateFunc, "template_function/*html")) // nah buat sekali saja di awal

type AnimeList struct {
	Title    string
	Realease string
}

func TemplateCacheing(writer http.ResponseWriter, request *http.Request) {
	anime := AnimeList{
		Title:    "Kaguya sama love is war seson 3",
		Realease: "April",
	}
	templateFuncCacheing.ExecuteTemplate(writer, "anime.html", anime)
}

func TestCacheing(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/kawai", nil)
	record := httptest.NewRecorder()

	TemplateCacheing(record, request)

	response := record.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
