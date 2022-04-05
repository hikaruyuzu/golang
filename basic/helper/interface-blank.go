package helper

// interface kosong dapat menangkap data apapun baik di parameter atau return value

// contoh function dengan return interface kosong
func FunctionBlankInterface(number int) interface{} {
	// jadi dia akan bisa mereturn data apapun itu
	if number == 1 {
		return "Hahh wakaranai"
	} else if number == 2 {
		return true
	} else {
		return 30
	}
}

// untuk datahasil pemanggilan function juga akan memiliki tipe data interface
