package helper

import (
	"os"
)

// nil adalah representasi data kosong seperti di bahasa pemrograman lain, seperti null
// nil, hanya dapat dipakai di funcion, interface, slice, map channel, dan pointer saja
// tipe data seperrti string, integer, bool tidak bisa menggunakan nil

// contoh penggunaan
func GetInfoEnv() interface{} {
	var info []string
	// mendapat info os
	var env = os.Environ()
	if env != nil {
		for _, envr := range env {
			info = append(info, envr)
		}
	}
	return info
}

// untuk pemanggilan kamu bisa gunakan seperti ini
/**
data := helper.GetInfoEnv()
// untuk pengecekan kosong atau tidak
if data != nil {
	fmt.Println(data)
} else {
	fmt.Println("Data kosong")
}
*/
