package helper

// cek palindrome
func Palindrome(value string) bool {
	var length = len(value)
	for i := 0; i < length/2; i++ {
		// value length-1-i mean slice start from last index - i
		head := value[i]
		tail := value[length-1-i]
		if head != tail {
			return false
		}

	}

	return true
}
