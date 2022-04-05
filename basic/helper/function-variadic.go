package helper

// function variadic adalah dimana kita bisa menginputkan lebih dari satu value nilai kedalam variable yang nantinya akan di ubah menjadi slice
// kamu bisa membuat variadic function di akhir dari function dengan simbol ...
// kamu hanya bisa membuatnya di final parameter yang artinya tidak bisa berada di depan
// jika kamu ingin memasukan value kedalam parameter misal sudah terlanjur dalam bentuk slice kamu bisa memakai nameSlice...
// menggunakan ... di belakang nama slice

// contoh pembuatan variadic function
func AddALlInputData(numbers ...int) int {
	var results = 0
	// disini dapat dilihat bahwa tipe dari numbers akan berubah menjadi slice
	for _, value := range numbers {
		results += value
	}
	return results
}
