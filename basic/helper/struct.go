package helper

import (
	"fmt"
	"strconv"
)

// struct berisi kumpulan data dari field dimana pada setiap field yang ada dalam struct kamu bisa membuatnya memiliki tipe data yang berbeda
// strcut sama halnya seperti class jika di bahasa berorientasi object
// struct biasanya digunakan untuk merepresentasikan tipe" data yang digunakakn untuk aplikasi yang kita buat
// cara pembuatannya jugaa ada banyak cara

// jika kamu ingin membuat struct dengan melibatkan pointer kamu bisa menggunakan keyword new(NameOfStruct)

// contoh cara pembuatan struct
type DataUserStruct struct {
	Username, Email string
	Age             int
}

// contoh pemanggilan struct
func AppendDataUser(name string, mail string, yourAge int) {
	// kita bisa membuat variable baru dengan tipe struct atau mendeklarasikan struct secara langsung
	// jika kamu ingin membuatnya menjadi pointer tinggal membuatnya menjadi car account = new(DataUserStruct)
	var accountUser DataUserStruct
	// kamu bisa menambah dan mengaksesnya dengan cara seperti ini
	accountUser.Username = name
	accountUser.Email = mail
	accountUser.Age = yourAge

	// kamu bisa akses dengan cara yang hampir sama
	fmt.Println(accountUser.Username)
	fmt.Println(accountUser.Email)
	fmt.Println(accountUser.Age)

}

// cara akses dnengan pendeklarasian secara langsung
func AppendDataUser2(name, mail string, yourAge int) interface{} {
	// nah ini adalah contoh pendeklarasian secara langsung
	accountUser2 := DataUserStruct{
		Username: name,
		Email:    mail,
		Age:      yourAge,
	}

	return accountUser2
}

// kamu juga bisa memasukan data kedalam struct tanpa harus mendeklarasikan nama fierldnya, namun cara ini kurang direkomendasikan
// karena rawanterjadinya error, dan harus mengisi semua nilai dari field, serta harus di isi secara berurutan
func DirectDeclareStruct() {
	accountUser3 := DataUserStruct{"Kato Megumi", "katochan@gmail.com", 18}
	fmt.Println(accountUser3)
}

// kamujuga bisa membuat struct di dalam struct dengan cara seperi halnya dibawah ini

// inner struct DataAddressCostumer
type DistrictSpesifik struct {
	District string
	Village  string
	Code     int
}

// inner struct Data Costumer
type DataAddressCostumer struct {
	Country  string
	Province string
	City     string
	District DistrictSpesifik
}

// maiin struct
type DataCostumer struct {
	Name    string
	Age     int
	Address DataAddressCostumer
}

func GetSpesifikDataUser() interface{} {
	// pemanggilan struct
	dataUser := DataCostumer{
		Name: "Kaguya Shinomiya",
		Age:  17,
		// call sub
		Address: DataAddressCostumer{
			Country:  "Japan",
			Province: "Tokyo",
			City:     "SHibuya",
			// call sub
			District: DistrictSpesifik{
				District: "SUb Shibuya",
				Village:  "Moekyun",
				Code:     61256,
			},
		},
	}

	return dataUser
}

// kamu juga bisa memasukan function kedalam struct acaranya sama seperti pada saat membuat function as parameter
// kita bisa memasukan functionnya kedalam type deklaration

// alias
type HelloMyWifeFromAnotherWorld func(name string) string

type HelloMyWIfe struct {
	name      string
	greetings HelloMyWifeFromAnotherWorld
}

// pemanggilan untuk struct HelloMyWIfe
func GetSayHelloFromIsekai(waifu string) string {
	mywife := HelloMyWIfe{
		name: waifu,
		greetings: func(name string) string {
			return fmt.Sprintf("Hello %s chan", name)
		},
	}
	return mywife.greetings(mywife.name)
}

// array of struct, kamu juga bisa menampung struct di dlam sebuah array/ slice
// struct untuk menampung nama user

type Username struct {
	name string
}

// function
func AppendUsername(test int) {
	// disini kamu harus membuat struct di dalam slice
	var tempUser []Username
	var results []Username

	for i := 0; i < test; i++ {
		data := strconv.Itoa(i)
		// add data to struct
		user := Username{data}
		results = append(tempUser, user)
		fmt.Println(results)
	}

}
