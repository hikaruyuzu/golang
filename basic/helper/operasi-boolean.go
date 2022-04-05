package helper

import (
	"fmt"
	"strconv"
)

// operasi boolean digunakan untuk membandingkan dua kondisi boolean atau dua atau lebih math operator
// yang akan menghasilkan nilai true atau false
// contoh beberpa operasi boolean
/*
	&& and , jika ada yang false di satu sisi maka dia akan bernilai false
	|| or, jika salah satu atau keduanya true maka dia akan true, sedangkan, jika semuanya false dia akan false
	! not misal !true maka dia akan menampilkan kebalikannya yaitu false, begitupun sebaliknya
*/

// kasus penggunaan sudah di terapkan di kasus- kasus sebelumnya
// contoh

// program menentukan kelulusan
func ProgramExam(value []int) string {
	var graduate string
	count := 0
	for i := 0; i < len(value); i++ {
		count += value[i]
	}

	average := count / len(value)
	averageString := strconv.Itoa(average)
	// disini adalah contoh penggunaanya
	if average >= 75 && len(value) >= 5 {
		graduate = fmt.Sprintf("Selamat kamu lulus dengan rata nilai %s", averageString)
	} else {
		graduate = "Kamu salah jurusan"
	}

	return graduate

}
