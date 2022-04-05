package helper

import "fmt"

// function parameter adalah cara pembuatan function dengan cara memasukan variable atau data yang dibutuhka dariluar function
// dengan menggunakan parameter setelah nama function dan dimasukan kedalam (parameter string) data tersebut(data yang di inputkan)
// nantinya daat diproses di dalam function yang sudah dibuat untuk tujuan tertentu
// kamu bisa membuat  banyak parameter didalam function

// contoh simple functio parameter
func SayDaisukiFromMyWaifu(name string, age int, from string) {
	// kamu bisa melaukan apa saja di block kode ini
	fmt.Println("Hello Shiyzu, my name is ", name, "aku dari anime ", from, "dan umurku sekarang ", age, "tahun, Daisuki ~ <3")
}
