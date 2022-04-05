package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

// digunakan untuk membuat json dengan tipe object
type Costumer struct {
	FirstName string
	LastName  string
	Age       int
	Married   bool
}

func logEncodeCostumer(data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestJSONEncode(t *testing.T) {
	costumer := Costumer{
		FirstName: "Rem",
		LastName:  "Chan",
		Age:       18,
		Married:   true,
	}
	logEncodeCostumer(costumer)
}
