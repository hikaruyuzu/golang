package helper

// di golang kita bisa mengenmbalikan value lebih dari 1 nilai
// kamu cukup menambahkan tipe datanya di belakang function sesuai dengan banyaknya data yang akan dikwmbalikan
// jika inginmengabaikan penangkapan value dari variablenya kamu cukup menggunakan _ saja

// contoh simple
func GetAllWaifu(name string, name2 string, name3 string) (string, string, string) { // nah string tiga kali iru adalah tipe return value
	return name, name2, name3
}
