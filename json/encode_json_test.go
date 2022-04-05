package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

// digunakan untuk encode dari tipe data golang ke json
func logJson(data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

type AnimeInfo struct {
	Title  string
	Rating int
}

func TestLogEncode(t *testing.T) {
	logJson("data")
	logJson(1)
	logJson(true)
	logJson([]string{"kaguya shinomiya", "kato megumi", "fuuka miyazawa"})
	logJson(AnimeInfo{
		Title:  "Kaguya sama love is war",
		Rating: 9,
	})
}
