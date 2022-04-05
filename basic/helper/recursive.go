package helper

// rekursve function adalah function yang memanggil drinya sendiri, yang akan terus berulang
// kita bisa menggunakan recursive dalam beberapa kasus misalnya kasus factorial

// ini adalah contoh kode recursive
func FactorialRecursive(number int) int {
	if number == 1 {
		return 1
	} else {
		return number * FactorialRecursive(number-1)
	}
}
