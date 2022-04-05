package helper

import "fmt"

// boolean direpresentasikan dengan bool yang bernilai true atau false

// function mencetak hallo
func HelloMyWaifu(name string) string {
	check := false
	if name != "" {
		// melakukan perubahan check dari false ke true jka kondisi terpenuhi
		check = true
	}

	var greeting string
	if check == true {
		greeting = fmt.Sprintf("Hello %s, my name is koichi", name)
	}
	return greeting
}

// functon untuk menentukan genre film
func AgeLimit(age int) string {
	var film string
	if age > 0 && age <= 13 {
		// disini akan menlakukan pengecekan perbandngan berniali true/ false jika true maka program akan di eksekusi jika false program akan mengecek kebawah, dan seterusnya
		film = "Genre comdey"
	} else if age > 13 && age < 18 {
		film = "Genre Romance, Comedy, Harem "
	} else {
		film = "Genre Romance, Comedy, Harem, Ecchi"
	}

	return film
}
