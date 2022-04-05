package helper

import "fmt"

// continue akan meng skip kode program di bawahnya jika kondisi yang ditentukan terpenuhi
// dengan kata lain dia akan melanjutkan ke eksekusi for loop selanjutnya

// contoh program untuk menentukan angka ganjil berikut
func GetOddValue(value int) {

	for i := 0; i < value; i++ {

		if i%2 == 0 {
			// nah disini jika kondisi terpenuhi dia akan langsung lanjut ke perulangan berikutnya daan akanmengabaikan kode di bawahnya
			continue
		}
		fmt.Println("Angka ganjil ke - ", i)
	}
}
