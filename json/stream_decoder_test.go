package json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

// pada kenyatannya kita akan sering mendapatkan input data request dari stream berupa io.reader, seperti request body, network, file, respnse code
// etc, nah ini sangat cocok menggunakan decoder stream daripada arsing dan mengubah tipe datanya secara manual, seperti emmasukan filenya secara mamnual
// kedalam variable terlebih dahulu

type DataVtuber struct {
	Name   string `json:"name"`
	Agensi string `json:"agensi"`
	Age    int    `json:"age"`
}

func TestStreamDecoder(t *testing.T) {
	reader, err := os.Open("data_vtuber.json")
	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(reader)

	dataVtuber := &DataVtuber{}
	if err := decoder.Decode(dataVtuber); err != nil {
		panic(err)
	}

	fmt.Println(dataVtuber)

}

func TestStreamEncoder(t *testing.T) {
	// digunakan secara langsung menulis response ke io.Writernya, tanpa harus melakukan konversi dahulu
	dataVTuber := DataVtuber{
		Name:   "kuroneko aii",
		Agensi: "hoshisho",
		Age:    17,
	}

	writer, err := os.Create("vtuber.json")
	if err != nil {
		panic(err)
	}

	encoder := json.NewEncoder(writer)

	err = encoder.Encode(dataVTuber)
	if err != nil {
		panic(err)
	}
}
