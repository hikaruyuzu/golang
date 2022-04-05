package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Country    string
	City       string
	PostalCode int
}

type Sellers struct {
	FullName  string
	Age       string
	Hobbies   []string
	Addresses []Address
}

func TestEncodeJSONSellers(t *testing.T) {
	data := Sellers{
		FullName: "Kato Megumi",
		Age:      "18",
		Hobbies:  []string{"Masak", "Nonton anime", "cooding"},
	}
	marshal, err := json.Marshal(data)
	if err != nil {
		return
	}

	fmt.Println(string(marshal))
}

func TestDecodeJSONSellers(t *testing.T) {
	jsonString := `{"FullName":"Kato Megumi","Age":"18","Hobbies":["Masak","Nonton anime","cooding"]}`
	jsonBytes := []byte(jsonString)

	sellers := &Sellers{}
	if err := json.Unmarshal(jsonBytes, sellers); err != nil {
		panic(err)
	}

	fmt.Println(sellers)
	fmt.Println(sellers.Hobbies)
}

func TestEncodeJSONSellersComplex(t *testing.T) {
	data := Sellers{
		FullName: "Kato Megumi",
		Age:      "18",
		Hobbies:  []string{"Masak", "Nonton anime", "cooding"},
		Addresses: []Address{
			{
				Country:    "Japan",
				City:       "Tokyo",
				PostalCode: 19218,
			},
			{
				Country:    "Japan",
				City:       "Shibuya",
				PostalCode: 29218,
			},
			{
				Country:    "Japan",
				City:       "Tokyo",
				PostalCode: 12318,
			},
		},
	}
	marshal, err := json.Marshal(data)
	if err != nil {
		return
	}

	fmt.Println(string(marshal))
}

func TestDecodeJSONSellersComplex(t *testing.T) {
	jsonString := `{"FullName":"Kato Megumi","Age":"18","Hobbies":["Masak","Nonton anime","cooding"],"Addresses":[{"Country":"Japan","City":"Tokyo","PostalCode":19218},{"Country":"Japan","City":"Shibuya","PostalCode":29218},{"Country":"Japan","City":"Tokyo","PostalCode":12318}]}`
	jsonBytes := []byte(jsonString)

	sellers := &Sellers{}
	if err := json.Unmarshal(jsonBytes, sellers); err != nil {
		panic(err)
	}

	fmt.Println(sellers)
	fmt.Println(sellers.Hobbies)
	fmt.Println(sellers.Addresses)
}

func TestEcodeOnlyArrayComplex(t *testing.T) {
	address := []Address{
		{
			Country:    "Japan",
			City:       "Tokyo",
			PostalCode: 19218,
		},
		{
			Country:    "Japan",
			City:       "Shibuya",
			PostalCode: 29218,
		},
		{
			Country:    "Japan",
			City:       "Tokyo",
			PostalCode: 12318,
		},
	}
	marshal, err := json.Marshal(address)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(marshal))
}

func TestDecodeOnlyArrayJSONSellersComplex(t *testing.T) {
	jsonString := `[{"Country":"Japan","City":"Tokyo","PostalCode":19218},{"Country":"Japan","City":"Shibuya","PostalCode":29218},{"Country":"Japan","City":"Tokyo","PostalCode":12318}]`
	jsonBytes := []byte(jsonString)

	address := &[]Address{}
	if err := json.Unmarshal(jsonBytes, address); err != nil {
		panic(err)
	}

	fmt.Println(address)
}

func TestDecodeOnlyObjectJSONSellersComplex(t *testing.T) {

	jsonString := `{"FullName":"Kato Megumi","Age":"18"}`
	jsonBytes := []byte(jsonString)

	sellers := &Sellers{}
	if err := json.Unmarshal(jsonBytes, sellers); err != nil {
		panic(err)
	}

	fmt.Println(sellers)
}
