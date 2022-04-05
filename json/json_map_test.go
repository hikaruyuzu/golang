package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

// json map digunakan jika data json tidak terprediksi misal dynmic, atau bisa berubah ubah baik jumlahnya
// ini akan lebih mudah unuk diguankana dariada json struct
// akan tetapi disini kita tidka bisa menggunakan tag di struct lagi

func TestJSONMapEncode(t *testing.T) {
	dataVtuber := map[string]interface{}{
		"name":   "shirayuki sora",
		"agensi": "hoshisho",
		"age":    18,
	}

	marshal, err := json.Marshal(dataVtuber)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(marshal))
}

// decode data json map
func TestJSONMapDecode(t *testing.T) {
	vtuberString := `{"age":18,"agensi":"hoshisho","name":"shirayuki sora"}`
	vtuberByte := []byte(vtuberString)

	var data map[string]interface{}

	err := json.Unmarshal(vtuberByte, &data)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
	fmt.Println(data["name"])
	fmt.Println(data["agensi"])
	fmt.Println(data["age"])

}
