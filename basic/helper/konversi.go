package helper

// digunakan untuk konversi tipe data dari satu tipe data ke tipe data lain
// kamu bisa menggunakan int32(namaVariableKamu)
// tapi disini kamu perlu berhati hati dikarenakan jika kamu mengkonversikan angka yang tidak memenuhi dia akan di looping kembali ke batas yang disediakan

// function untuk konversi float ke int
func AgeKonvertion(age float32) int8 {
	ageNow := age
	// konversi ke int8
	var ageInInteger int8 = int8(ageNow)
	return ageInInteger
}

// konversi byte  atau alias dari uimt8 ke string
func GetALetlers(kata string) int64 {

	countOfA := 0
	for i := 0; i < len(kata); i++ {
		// disini saya melakukan konversi byte ke string
		if string(kata[i]) == "a" || string(kata[i]) == "A" {
			countOfA += 1
		}

	}
	return int64(countOfA)

}
