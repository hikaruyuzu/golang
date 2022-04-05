package gopackage

import (
	"fmt"
	"strings"
)

// package string digunakan untuk memdifikasi string
// ada banyak sekali yang bisa kita lakukan untuk memanipulasi string
// kamu bisa baca selengkapnya di https://pkg.go.dev/strings

func ModificationString() {
	// untuk mencari nilai dari string
	fmt.Println(strings.Contains("Kaguya sama: love is war", "love")) // true
	fmt.Println(strings.Contains("Kaguya sama: love is war", "chan")) // false

	// untuk memotong string menjadi slice
	fmt.Println(strings.Split("Kaguyya sama: Love is War", " "))

	// untuk memtong kata/ symbol, etc berdasarkan separator, akan menghilangkan %
	fmt.Println(strings.Trim("%kaguya sama: love is war%", "%"))

	// to uupercase
	fmt.Println(strings.ToUpper("kaguya sama love is war"))
	// to lowercase
	fmt.Println(strings.ToLower("KAGUYA SAMA LOVE IS WAR"))

	// mengubah nilai
	fmt.Println(strings.ReplaceAll("kaguya kaguya kaguya", "kaguya", "megumi"))

}
