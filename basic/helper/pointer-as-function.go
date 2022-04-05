package helper

// kamu bisa menggunakan pointer di uncttion untuk memanipulasi data diluar function
// jika tidak menggunakan function maka data yang di ubah itu hanya data yang ada di dalam function saja yang merupakan data copy
// buka data aslinya jadi jika ingin manipulasi data asli kamu bisa menggunakan pointer sebagai paramater
// denga cara menambahkan tanda * di paramater

// struct data

type AnimeInfo struct {
	Title   string
	Studio  string
	Country string
}

// contoh jika tidak menggunakan pointer
func ChangeStudioAnimeNoPointer(studio AnimeInfo, newStudio string) {
	// dia tidak akan berubah walau functionnya dipanggil
	studio.Studio = newStudio
}

// contoh function yang menggunakan pointer
// nah disini kamu harus mengunakan tanda *, untuk memberi tah bahwa dia harus parameter pointer
func ChangeStudioANimeWithPointer(studio *AnimeInfo, newStudio string) {
	// dia akan mengubah main data menjadi newStudio yang d inputkan
	studio.Studio = newStudio
}

/**
/ pemanggilan

// data awal
anime := helper.AnimeInfo{
	Title:   "Eromanga sensei",
	Studio:  "WIT-Studio",
	Country: "Japan",
}
// disini nama studionya akan sama saja, karena tidak menggunakank pointer
helper.ChangeStudioAnimeNoPointer(anime, "A1-Picture")
fmt.Println(anime)

// ini yang menggunakan poninter akan mengubah nama studionya
// akan tetapi kamu harus membuat parameter yang harus emnggunakan pointer dengan key, &, yang kaan menjadikannya pass by reference

// atau kamu bisa membuat parameter sebelumya yang merupakan pointer
var animePointer *helper.AnimeInfo = &anime
helper.ChangeStudioANimeWithPointer(animePointer, "A1-Picture")
fmt.Println(anime)

*/
