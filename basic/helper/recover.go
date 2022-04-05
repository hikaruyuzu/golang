package helper

import "fmt"

// recover adalah function untuk menangkap panic agar program tetap berjalan walaupun terjadi error
// kamu bisa membuat recover ini di dalam kode defer agar progarm panic dapat ditanngkap oleh recover sebelum program berhenti

func EndMyApp() {
	// disini recover akan menangkap error panic 
	getPanicMessage := recover()
	if getPanicMessage != nil {
		fmt.Println("error with message, ", getPanicMessage)
	}
	fmt.Println("APPLIKASI SELESESAI DI JALANKAN")
}

func StartAplicationRecover(konvirmation bool) {
	// walau error akan tetap di eksekusi
	defer EndMyApp()
	if konvirmation {
		panic("DISINI ERROR DONGG")
	}

	fmt.Println("APLIKASI SELESAI DIJALANKAN")
}
