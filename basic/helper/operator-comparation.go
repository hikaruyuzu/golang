package helper

// operasi perbandingan dapat diguakan untuk membandingkan tipe data number ataupun string
// hasil perbandingan dari operasi perbandingan dalah boolean baik boolean true ataypun boolean false
// jenis operasi perbandingan yang bisa kamu gunakan

/*
>
<
>=
<=
!=
==
*/

// untuk cara penggunaanya kamu bisa lihat di file boolean.go atau di file lainnya yang sudah di buat
// atau kamu bisa melihat contoh kasus palindrome brikut ini

// function palindrome
func PalndromeCheckLetters(letters string) bool {
	palindrome := true
	l := len(letters)
	// disini kasus penggunaan operasi perbandingan
	for i := 0; i < l/2; i++ {

		head := letters[i]
		// check dari index terakhir - i
		tail := letters[l-1-i]

		// contoh perbandingan ==
		if head != tail {
			palindrome = false
		}
	}
	return palindrome
}
