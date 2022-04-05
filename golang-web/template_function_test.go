package golangweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// ini digunakan untuk membuat functon di dalam file hmtl
// template function ada beberaoa macam yaitu
// template function biasa, template function global(bisa di akses secara langsung seperti template function operasi perbandingan/ kamu
// bisa lihat dokumentasinya di google ada banyak, dan kamu juga bisa membuatnya sendiri), dan template pipeline yaitu megirim hasil
// function dari pemanggilan pertama ke pemanggilan function berikutnya dengan menggunakan tanda |

type Anime struct {
	Title    string
	Realease string
}

func (anime Anime) AnimeInfo(character string) string {
	return fmt.Sprintf("Anime dengan judul %s dengan main character %s akan tayang pada bulan %s", anime.Title, character, anime.Realease)
}

func TemplateFunction(writer http.ResponseWriter, request *http.Request) {
	// ke file htmlnya
	t, err := template.ParseFiles("./template_function/anime.html")
	if err != nil {
		panic(err)

	}

	anime := Anime{
		Title:    "Kaguya sama love is war seson 3",
		Realease: "April",
	}
	t.ExecuteTemplate(writer, "anime.html", anime)
}

func TestTemplateFunctionLive(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/anime", TemplateFunction)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestTemplateFunction(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/kawai", nil)
	record := httptest.NewRecorder()

	TemplateFunction(record, request)

	response := record.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

// template function global
// https://github.com/golang/go/blob/master/src/text/template/funcs.go

func TemplateFunctionGlobal(writer http.ResponseWriter, request *http.Request) {
	// len adalah global function dari dokumentasi github di atas
	t, err := template.New("FUNCTION").Parse(`{{ .AnimeInfo "Kaguya Shinomiya" | len}}`) // | adalah pipeline
	if err != nil {
		panic(err)
	}
	anime := Anime{
		Title:    "Kaguya sama love is war seson 3",
		Realease: "April",
	}

	t.ExecuteTemplate(writer, "FUNCTION", anime)

}

func TestTemplateFunctionGlobal(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/kawai", nil)
	record := httptest.NewRecorder()

	TemplateFunctionGlobal(record, request)

	response := record.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

// nah bagaimana caranya membuat function global sendiri ?
func CreateFunctionGlobal(writer http.ResponseWriter, request *http.Request) {
	t := template.New("FUNCTION")
	// mendaftarkan global func
	t.Funcs(map[string]interface{}{
		"upper": func(name string) string {
			return strings.ToUpper(name)
		},
	})

	// | adalah pipline dibawah ini artinya memanggil function AnimeInfo dengan paramerter kaguya shinomiya
	// lalu hasilnya akan di masukan ke function berikutnya untuk di upper
	tmpl, err := t.Parse(`{{.AnimeInfo "kaguya shinomiya" | upper }}`) // | adalah pipeline
	if err != nil {
		panic(err)
	}

	anime := Anime{
		Title:    "Kaguya sama love is war seson 3",
		Realease: "April",
	}

	tmpl.ExecuteTemplate(writer, "FUNCTION", anime)
}

func TestTemplateFunctionGlobalSolo(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/kawai", nil)
	record := httptest.NewRecorder()

	CreateFunctionGlobal(record, request)

	response := record.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

// | adalah pipeline
