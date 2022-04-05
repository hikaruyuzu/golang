package golangweb

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// upload file menggunakan multipart form
//go:embed form_upload
var formUpload embed.FS

// kita buat cacheing agar tidak lemot
var formUploadEmbed = template.Must(template.ParseFS(formUpload, "form_upload/*.html"))

// handler render file
func UploadFileRender(writer http.ResponseWriter, request *http.Request) {
	formUploadEmbed.ExecuteTemplate(writer, "upload.html", nil)
}

// handler upload file
func UploadFileForm(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseMultipartForm(100 << 20) //  max 100mb
	if err != nil {
		panic(err)
	}
	// membaca body dan hasil parsing
	file, fileHeader, err := request.FormFile("file")
	if err != nil {
		panic(err)
	}
	// membuat tempat menaruh hasil file upload
	destination, err := os.Create("./form_upload/" + fileHeader.Filename)
	if err != nil {
		panic(err)
	}
	// copy file upload ke destination
	_, err = io.Copy(destination, file)
	if err != nil {
		panic(err)
	}
	// untuk handle request post name
	err = request.ParseForm()
	if err != nil {
		panic(err)
	}
	name := request.PostForm.Get("textname")
	formUploadEmbed.ExecuteTemplate(writer, "success.html", map[string]interface{}{
		"Name": name,
		"File": "/static/" + fileHeader.Filename,
	})
}

func TestUploadFileServer(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", UploadFileRender)
	mux.HandleFunc("/upload", UploadFileForm)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./form_upload"))))
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// //go:embed form_upload/moetachan3.jpg
var imageTest []byte

func TestUploadFile(t *testing.T) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	writer.WriteField("textname", "kaguya chan")

	file, err := writer.CreateFormFile("file", "moeta")
	if err != nil {
		panic(err)
	}
	file.Write(imageTest)
	writer.Close()

	request := httptest.NewRequest(http.MethodPost, "http://localhost/test", body)
	// header sudah menjadi standard
	request.Header.Add("content-type", writer.FormDataContentType())
	record := httptest.NewRecorder()

	UploadFileForm(record, request)

	response := record.Result()
	bodyResponse, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bodyResponse))
}
