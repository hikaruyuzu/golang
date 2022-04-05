package embed

import (
	"embed"
	"fmt"
	"io/ioutil"
	"testing"
)

// embed multiple file di golang

//go:embed files/xut1.txt
//go:embed files/xut2.txt
//go:embed files/xut3.txt

var marinchan embed.FS // embed.FS akan menampung smeua hasil embed, dia adalah tipe data embed

func TestReadMultipleEmbed(t *testing.T) {
	aw, _ := marinchan.ReadFile("files/xut1.txt")
	fmt.Println(string(aw))

	aww, _ := marinchan.ReadFile("files/xut2.txt")
	fmt.Println(string(aww))

	awww, _ := marinchan.ReadFile("files/xut3.txt")
	fmt.Println(string(awww))
}

// nah daripadata memasukan satu persatu di dalam embednya lebih enak menggunakan path matcher
// baca di package path matcher golang doc, contoh path matcher seperti *, ? etc

//go:embed files/*.txt
var sinopsis embed.FS

func TestPathMatcher(t *testing.T) {
	dirEntry, _ := sinopsis.ReadDir("files")
	for _, entry := range dirEntry {
		if !entry.IsDir() {
			fileName := entry.Name()
			fmt.Println(fileName)
			body, _ := ioutil.ReadFile("files/" + fileName)
			fmt.Println(string(body))
		}

	}
}
