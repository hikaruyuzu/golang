package helper

// array adalah tpe data yang dapat menampung banyak data dengan satu tipe yang sudah di tentukan panjangnya sejak awal
// yang dapat di akses dengan menggunakan no index 0 - panjang array-1
// kamu bisa membuat array di golang dengan cara seperti ini

var data [10]string // ini adalah array kosong
// kalu kamu ingin mengisi data dari si array kamu bisa menggunakan nama variable[nomor indexny]
// misal menambahkan data ke varaible data

// data[0] = "Kaguya sama"
// data[1] = "Kaguya sama"
// data[2] = "Kaguya sama"

// ini adalah contohb jika ingon mendeklarasikanya secara langsung
var listName = []string{"Ranko Kanzaki", "Kato Megumi", "Tsukasa Chan", "Mai Sakurajima"}

// untuk akses juga sama aja kamu bisa menggunakan no index yang ada
// fmt.Println(data[0])

// untuk mendapatkan panjangnya kamu bisa memakai function len(data) yang ada di array

// contoh program nemabhakan angka bilangan genap dari 0 - angka yang di inputkan
func AddDataEven(data int) []int { // return data berupa array
	// membuat data array dengan tampungan maximal 1000 data
	var even [1000]int
	array := 0
	// looping data
	for i := 0; i <= data; i++ {
		if i%2 == 0 {
			// menambahkan data genap ke dalam array
			even[array] = i
			array += 1
		}
	}
	return even[:] // return semua data
}
