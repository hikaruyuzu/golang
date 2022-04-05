package helper

import "fmt"

// panic ini digunakan untuk memberi result pada error yang nantinya dapat di tagkap oleh recover

func StartPanic(result bool) {
	if result {
		panic("ERROR NIHH") // jika panic akan di costum dengan message di dalam panic yang dibuat
	}
	fmt.Println("JIKA ERROR PROGRAM DISINI AKAN BERHENTI") // gabakal di eksekusi kalau panic
}
