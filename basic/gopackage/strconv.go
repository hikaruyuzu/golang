package gopackage

import (
	"fmt"
	"strconv"
)

// digunkana untuk konversi tipe data
// lebih lengkapnya kamu bisa baca di https://pkg.go.dev/strconv

func ChangeDataTYpeStrconv() {
	// yang di maksud base disini adalah base dari misal desimal berarti 10, boolean= 2, hexa = 16, octal = 8

	dataInt, err := strconv.ParseInt("1000000", 10, 64)
	if err != nil {
		fmt.Println("error ", err.Error())
	} else {
		fmt.Println(dataInt)
	}

	// kalau dari some data type ke string kamu bisa memakai Format
	dataString := strconv.FormatBool(true)
	fmt.Println(dataString)

	// base 2 berarti dia akan menjadi binner
	dataNumberString := strconv.FormatInt(1000000, 2)
	fmt.Println(dataNumberString)

	// adapun cara simple yaitu menggunakan Atoi(ke int) dan Itoa(ke string)

}
