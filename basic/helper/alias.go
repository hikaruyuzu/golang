package helper

import (
	"fmt"
	"strconv"
)

// alias digunakan untuk membuat type data baru dari tipe data yang sudah ada
// kamu bisa menggunakan keyword type untuk membuatnya

// function untuk memasukan data diri

// ini adalah alias
type noKtpUser string

func CreateMyData(user noKtpUser, noktp string) string {

	var no noKtpUser = user
	// harrus dikonversi walau menggunakan alias
	var noKtp noKtpUser = noKtpUser(noktp)

	return fmt.Sprintf("My name is %s with no %s", string(no), string(noKtp))
}

// contoh lain yang simple
type Data string

var DataUser Data = "Kaguya Shinomiya"

func GetAge(age int) Data {
	myAge := age
	myAgeDataString := strconv.Itoa(myAge)

	return Data(myAgeDataString)
}
