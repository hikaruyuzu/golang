package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

// json tag digunakan untuk menyesuaikan carapenulisan best practice field di json dan field struct atau attribute lainnya
// agar tetap bisa mengikuti best practice masing-masing

type AnimeDetails struct {
	Title    string `json:"title"`
	Rating   string `json:"rating"`
	Studio   string `json:"studio"`
	ImageURL string `json:"image_url"`
}

func TestEncodeJSONAnime(t *testing.T) {
	anime := AnimeDetails{
		Title:    "RE:ZERO",
		Rating:   "8",
		Studio:   "White-Fox",
		ImageURL: "http://images/rezero.png",
	}

	marshal, err := json.Marshal(anime)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(marshal))
}

// decode
func TestDecodeJSONAnime(t *testing.T) {
	jsonString := `{"title":"RE:ZERO","rating":"8","studio":"White-Fox","image_url":"http://images/rezero.png"}`
	jsonByte := []byte(jsonString)

	animeDetails := &AnimeDetails{}
	if err := json.Unmarshal(jsonByte, animeDetails); err != nil {
		panic(err)
	}
	fmt.Println(animeDetails)
}
