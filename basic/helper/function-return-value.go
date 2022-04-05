package helper

import "fmt"

// function return value adlah dimana function yang dibut dapat mengembalikan data sesuai tipe data yang kita inginkan,
// yang nantinya bisa di tangkap oleh main program atau function lain

// misal saya ingon mengembalikan tipe data berupa string, maka di akhir pendeklarasian function harus ditulis tipe datanya apa

// contoh function return value sederhana
func GetYourWaifu(name string) string { // nah string di akhir merupakan pendeklarasian return value, ini bisa di isi tipe data apa saja
	// disini kita bisa melakukan operasi apapun
	if name != "" {
		return fmt.Sprintf("i love you %s ", name)
	} else {
		return fmt.Sprintf("Hii nama kamu siapa? ")
	}
	// block kode apapun setelah retun value tidak akan di eksekusi lagi
}
