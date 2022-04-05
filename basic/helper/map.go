package helper

import (
	"fmt"
	"strconv"
)

// map memilik dua komponen utama yatu key - value
// dimana key dapat berupa string ataupun number, yang bisa di tentukan secara random asal tidak sama
// jika key nya sama maka dia akan mengantikan data sebelumnya yang di inputkan

/*
/ 	keyword untuk membuat map
	// map kosong
	var nameMap = make(map[string]string)  -> untuk []pertama itu adalah nama tipe data dari keyvaluenya dan belakangnya adalah valuenya

	// mendapatkan jumlah data dalam map
	len(nameVariableMap)

	// menghapus data dalam map
	delete(namaVariableMap, "keynya")

	// menambahkan data dalam map
	namaVaraible["keybaru"] ="valuenya"

	// map dengan isi default atau membuat map langsung
	datauser := map[string]string {
		"name" : "kaguya shinomiya"
		"anime" : "kaguya sama love is war"
	}

*/

// contoh perogram dengan mengguankan map
func GetDataUserMap(name string, email string, address string) interface{} {
	// membuat map baru
	var dataUser = make(map[string]string)
	// menambahkan data kedalam slice
	dataUser["name"] = name
	dataUser["email"] = email
	dataUser["address"] = address

	if dataUser["address"] == "japan" {
		delete(dataUser, "email")
	}

	return dataUser
}

// jika input data secara langsung kamu bisa buat sepertiini
var dataWaifu = map[string]string{
	"name":  "Kato Megumi",
	"email": "katomeguchan@gmail.com",
	"age":   "19",
}

func DefaultUser() {
	fmt.Println(dataWaifu["name"])
	fmt.Println(dataWaifu["email"])
	age := dataWaifu["age"]

	ageInt, _ := strconv.Atoi(age)
	fmt.Println(ageInt)

}
