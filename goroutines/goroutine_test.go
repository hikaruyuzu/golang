package goroutines

import (
	"fmt"
	"testing"
	"time"
)

// kamu bisa membuat goroutines dengan menambhakan keyword go di depannama function
// go funstionName

// contoh
func HelloWorld(name string) {
	fmt.Println("Hello ", name)
}

func TestHelloWorld(t *testing.T) {
	// jika tidak menggunakan goroutines maka program akan berjalan secara parallel, dan tidak ascyncronus
	// yang artinya dia akan menunggu sampai function HelloWorld selesai di eksekusi

	// jika menggunakan goroutines dia akan secara otomatis mengeksekusi progra selanjutnya tanpa menunggu funcion selesai di eksekusi
	// time digunakan unuk menunggu
	// tambahkan go untuk membuatnya berjalan sebagai goroutines

	go HelloWorld("Kaguya shinomiya")
	// firs execute
	fmt.Println("Oppsss")
	// menunggu goroutines selesai
	time.Sleep(1 * time.Millisecond)
}

// goroutine sangat" ringan dan murah
// kamu hanya memperlukan 2kb untuk / goroutines
// mari kita coba dengan membuat 1000k goroutines

func SayfromGoroutine(number int) {
	fmt.Println("number ", number)
}

func TestSayFromGoroutine(t *testing.T) {
	for i := 0; i < 1000000; i++ {
		go SayfromGoroutine(i)
	}
	time.Sleep(10 * time.Second)
}
