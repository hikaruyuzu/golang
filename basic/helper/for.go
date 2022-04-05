package helper

import (
	"fmt"
)

// komonen for loop terdiri dari 3 komponen yaitu

// for initStatement; kondisiStatemrnt; postStatement {
// 		blockkode
// }

/*
init statement berisi kondisi awal dimana perulangan akan dimulai
kondisi berisi, persyaratan kondisi yang harus dipenuhi
post statement berisi penambahan kondisi awal dari setiap kali perulangan

/	siklus kerja
	init statement -> cekKOndisi -> kodeProgramDiEksekusi -> post statement -> cekKondisi -> kodeProgramDiEskesusi -> postStatement
*/

// conmtoh func mencetak agka dari 0 sampai yang di inputkan
func PrintAngka(number int) {
	// i adalah nit statement
	// i < number adalah kondisi
	// i++ adalah post statement
	for i := 0; i < number; i++ {
		fmt.Println("Angka, ", i)
	}
}

// kamu juga bisa memanfaatkan for loop untuk pengaksesan dan modifikasi index pada array atau slice

// kitabisa menggunakan for range untuk alternatif lain dalam pengaksesan array atau map

// func mendapatkan nama waifu
func GetNameWaifu() {
	var waifu = make(map[string]string)
	waifu["name1"] = "kaguya shinomiya"
	waifu["name2"] = "Kato megumi"
	waifu["name3"] = "Chino kaffu"
	waifu["name4"] = "Tsukasa Chan"
	waifu["name5"] = "Ranko Kanzaki"

	for key, value := range waifu {
		fmt.Println(key, " = ", value)
	}
}
// kamu bisa menggunakan tanda _, untuk mengganti key/ value jika tidak dipakai