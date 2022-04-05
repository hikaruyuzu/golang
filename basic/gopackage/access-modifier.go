package gopackage

import "fmt"

// akses modifier adalah suatu cara untuk menentukan apakah suatu variable boleh di akses dari luar package atau tidak
// disini caranya cukup simple tidak usah menggunakan kata kunci bantu
// kita cukup menggunakan huruf awalana kecil untuk membuatnya menjadi private
// atau membuatnya dengan awlaan huruf besar jika ingin membuatnya menjadi publik

// contoh jika di variable, ini dapat di akses dari luar package
var WaifuVersion = "waifu 1.0.0"

// ini akan error
// var versionWaifu = "waifu 1.0.1"

// jika menggunakan function, ini contoh jika private
func sayNyanpasu() {
	fmt.Println("Nyanpasu ....")
}

// ini contoh untuk publik dapat di akses dari luar package
func SayNyanpasu(name string) {
	fmt.Println("Nyanpasu ", name)
}
