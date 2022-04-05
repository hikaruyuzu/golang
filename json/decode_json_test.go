package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type DetailsAnime struct {
	Title  string
	Studio string
	Rating int
}

// encode json
func EncodeDetailsAnime(data interface{}) []byte {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	return bytes
}

// decode json
func DecodeJSON(dataJSON []byte, detailsAnime DetailsAnime) *DetailsAnime {

	jsonDetailsAnimeDecode := &detailsAnime
	err := json.Unmarshal(dataJSON, jsonDetailsAnimeDecode)
	if err != nil {
		panic(err)
	}

	return jsonDetailsAnimeDecode
}

func TestDecodeJSON(t *testing.T) {
	animeBefore := DetailsAnime{
		Title:  "Re:zero",
		Studio: "white-fox",
		Rating: 9,
	}
	// encode untuk di decode
	encodeAnime := EncodeDetailsAnime(animeBefore)
	fmt.Println(string(encodeAnime))

	// sesudah decode
	dataDecode := DecodeJSON(encodeAnime, DetailsAnime{})
	fmt.Println(dataDecode)
	fmt.Println(dataDecode.Title)
	fmt.Println(dataDecode.Studio)
	fmt.Println(dataDecode.Rating)
}

// simple version
func TestDecodeSimpleToUnderstand(t *testing.T) {
	jsonString := `{"Title":"Re:zero","Studio":"white-fox","Rating":9}`
	jsonByte := []byte(jsonString)

	jsonDecode := &DetailsAnime{}
	// konversi, otomatis hasil decode akan dimasukan ke parameter ke dua
	if err := json.Unmarshal(jsonByte, jsonDecode); err != nil {
		panic(err)
	}

	fmt.Println(jsonDecode)
	fmt.Println(jsonDecode.Title)
	fmt.Println(jsonDecode.Studio)
	fmt.Println(jsonDecode.Rating)

}
