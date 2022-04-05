package helper

import "fmt"

// defer digunakan agar programtetap di eksekusi walau progrqam terdapat error, defer ini akan sering digunakan

func EndApplication() {
	fmt.Println("APPLIKASI SELESESAI DI JALANKAN")
}

func StartApplicationDefer(number int) int {
	// walau error ini akan tetap dijankan
	defer EndApplication()
	result := 10 / number
	// jika program terdapat error ini tidak akan dijalankan
	fmt.Println("Aplikasi selesai dijalankan")
	return result
}
