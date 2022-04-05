package gopackage

import (
	"fmt"
	"reflect"
)

// package reflection, biasanya digunakan untuk melihat struktur kode kita saat aplikasi sedang berjalan
// refflect sangat berguna ketika kita ingin membuat library yang general
// lebih lengkapnya kamu bisa baca di sini https://pkg.go.dev/reflect

// example
type UserAccount struct {
	// ini adalah struct tag, yang nantinya dapat digunakan nilainya, untuk malakukan berbagai hal termasuk validation
	Name  string `required:"true" max:"24"`
	Email string `required:"true" have:"@"`
}

func GetReflectionTag(data interface{}) {
	dataType := reflect.TypeOf(data)
	// length
	fieldLength := dataType.NumField()
	fmt.Println(fieldLength)
	// mendapatkan nama fiekd berdasarkan index
	dataField1 := dataType.Field(0)
	dataField2 := dataType.Field(1)

	fmt.Println(dataField1.Name)
	fmt.Println(dataField2)

}

// dari sini kamku bisa membuat valiadasi simple dengan struct tag, dari fitur reflection
func IsValid(data interface{}) bool {
	typeField := reflect.TypeOf(data)
	// get len
	lengthField := typeField.NumField()
	for i := 0; i < lengthField; i++ {
		// get required
		field := typeField.Field(i)
		required := field.Tag.Get("required")
		// have := field.Tag.Get("have")
		if required == "true" {
			// cek empty
			value := reflect.ValueOf(data).Field(i).Interface()
			if value == "" {
				return false
			}
		}
	}
	return true
}
