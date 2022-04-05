package helper

import (
	"fmt"
	"strconv"
)

// if merupakan percabangan dimana jika kondisi terpenuhi maka perogram akan di jalankan, jika kondisi tidak terpenuhi dia akan mengecek ke
// konidisi di bawahnya
// adapaun selain if yaitu else if, dan else

func GetSayName(yourName string) string {
	var greetings string
	if yourName == "kaguya shinomiya" {
		greetings = fmt.Sprintf("Hello %s ayangku", yourName)
	} else if yourName != "" {
		greetings = fmt.Sprintf("Hello %s", yourName)
	} else {
		greetings = "Hello kenalan yukk!"
	}

	return greetings
}

// adapunversi singkat dari if statement yaitu misal untuk mengukur panjang dari kata yang ada
func GetLengthOfKata(yourLetters string) string {
	var length string
	if l := len(yourLetters); l <= 6 {
		length = fmt.Sprintf("Sorry, nama anda terlalu singkat yaitu, %s  kata", strconv.Itoa(len(yourLetters)))
	} else {
		length = fmt.Sprintf("Nama kamu %s kata, good", strconv.Itoa(len(yourLetters)))
	}
	return length
}
