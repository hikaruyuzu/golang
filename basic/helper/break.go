package helper

import "fmt"

// break digunakan untuk menghentikan perulangan di sata kondisi yang di tentukan terpenuhi

// func akan menghentikan perulangan saat kondisi value == 6
func BreakIfSix(value int) {
	for i := 0; i < value; i++ {

		if i == 6 {
			// nah jika kondisi di atas terpenuhi maka kode akan dihentikan, dan kode di bawahnya tidak akan di eksekusi
			break
		}

		fmt.Println("value ke - ", i+1)
	}
}
