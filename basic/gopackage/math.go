package gopackage

import (
	"fmt"
	"math"
)

// math berisi kumpulan konstan dan fungsi untuk math
// kamu bisa melihat dokumentasinya di https://pkg.go.dev/math

func MathGolang() {
	// round membulatkan ke yang paling dekat
	fmt.Println(math.Round(1.8))
	fmt.Println(math.Round(1.3))

	// ceil digunakan untuk membulatkan ke atas
	fmt.Println(math.Ceil(1.2))
	// floor digunakan untuk membulatkan ke bawah
	fmt.Println(math.Floor(1.9))

	// max digunakan untuk membandingkan yang terbesar
	// min digunakan untuk membandingkan yang terkecil
	// cara penggunaan sama saja

}
