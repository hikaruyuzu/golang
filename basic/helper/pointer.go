package helper

import "fmt"

// secara default golang menggunakan konsep pass by value
// yang artinya jika datanya akan dicopy, tidak pas by reference atau mengacu pada sumber yang sama

// dengan pointer kita bisa membuat datanya menjadi pass by reference atau mengacu pada sumbber yang sama
// agar tidak terjadi duplikasi data
// kamu bisa membuatnya dengan keyword &, jika ingin membuatnya menjadi pointer
// akan tetapi dia hanya akan merubah data subnya, jika inginmengubah keseluruhan data secara bulat"
// kamu bisa menggunakan keyword *diDepanVaribbaleBaru, ini akan merubah semua isi dari variable yang referene ke dia

// jika kamu ingoin membuat sebuah pointer baru(kosong) kamu bisa menggunakan keyword new(nameOfStruct)

// example
type CostumerAddress struct {
	Name                    string
	City, Province, Country string
}

// function pointer
func LearnPointerPassByValue() {
	user := CostumerAddress{"Kaguya Shinomiya", "Semarang", "Jawa Tengah", "Indonesia"}
	// disini dia akan pass by value, jika di ubah user1 maka user datanya tidak akan berubah
	user1 := user
	// misal kita ubah yaa
	user1.Name = "Kato Megumi"
	user1.City = "Pati"

	fmt.Println(user)  // data user
	fmt.Println(user1) // data user 1
}

// kalau dengan pointer(pass by reference datanya akan berubah)
func LearnPOinterPassByReference() {
	user := CostumerAddress{"Kaguya Shinomiya", "Bandung", "Jawa Barat", "Indonesia"}

	// kamu bisa membuat pass by reference dengan keyword &
	// disini artinya user 1 adalah pointer dari user, yang artinya jika data user1 di ubah maka data user akan ikutberubah
	var user1 *CostumerAddress = &user
	// di ubah misal
	user1.Name = "Ranko Kanzaki"
	user1.City = "Jakarta"

	fmt.Println(user) // nah disini datanya akan ikut seperti data user1(akan berubah)
	fmt.Println(user1)

}

// nah kalau kamu ingin emngubah datanya bulat"(semuanya), kamu tidak bisa menggunakan keyword &, itu akan menyebabkan datanya di copy
// kamu harus memggunakan keyword *sebelum nama dari variable

func LeranPointerChangeAllData() {
	// misal data awal ini
	user := CostumerAddress{"Kaguya Shinomiya", "Bandung", "Jawa Barat", "Indonesia"}

	// nah ini gabisa yaa
	// user1 := &user
	// user1 := &CostumerAddress{"Kato megumi", "Cirebon", "Jawa Barat", "Indonesia"}

	// kamu haru membuatnys menggunakan *, seperti berikut
	// jadi nanti misal ada variable yang reference ke struct yang sama mau berapapun itu, nilainya akan tetap sama
	user1 := &user
	*user1 = CostumerAddress{"Kato megumi", "Cirebon", "Jawa Barat", "Indonesia"}

	// misal reference lagi ke user, maka datanya akan tetap sama saja
	user2 := &user

	fmt.Println(user)
	fmt.Println(user1)
	fmt.Println(user2)

}

// misal kamu ingin membuat struct pointer kosong kamu bisa gunakan keyword new(nameOfStruct)
func LearnPointerCreateNewUser(name, city, province, country string) *CostumerAddress {
	// membuat empty user
	user := new(CostumerAddress)

	// mengisi field dengan isi parameter
	user.Name = name
	user.City = city
	user.Province = province
	user.Country = country

	return user
}
