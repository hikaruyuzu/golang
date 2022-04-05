package helper

// string direpresentasikan dengam namavariable di ikuti dengan string contoh name string

// function unruk mengambil panjang dari string
func GetLength(letters string) int32 {
	count := len(letters)
	return int32(count)
}

// function palindrome
func PalndromeCheck(letters string) bool {
	palindrome := true
	l := len(letters)
	for i := 0; i < l/2; i++ {

		head := letters[i]
		// check dari index terakhir - i
		tail := letters[l-1-i]

		if head != tail {
			palindrome = false
		}
	}
	return palindrome
}
