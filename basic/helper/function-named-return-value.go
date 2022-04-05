package helper

// function named return value hanya ada di golang saja sampai saat ini 31 januari 2022
// kamu bisa membuatnya sepertihalnya saat kamu membuat nama dari paramater di function dengan namaVariable string(diikuti dengan tipe datanya)
// misal

func GetFullNameWaifu(firstName string, middleName string, lastName string) (first string, middle string, last string) {
	// kamu bisa langsung mendeklarasikannya dengan nama returnnya
	first = firstName
	middle = middleName
	last = lastName

	// dan kamu tidak perklu repot' menulis semua nama retrunnya tapi cukup dengan menggunakan key return saja
	return
}
