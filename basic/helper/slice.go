package helper

import "fmt"

// disini kita dapat menympulkan bahwa slice adalah merupakan potongan dari array
// jika nilai pada array di ubah maka nilai dari slice juga akan berubah

// di dalam slice ada bagiaan yang dinamakan pointer, yaitu bagian terdepan dari slice, atau dimana slice dimulai slice{4:6} berarti slice di
// awali pada index ke 4 dan di akhiri pada index 6-1 = yaitu 5

// ada juga bagian yang dinamakan length yaitu seberapa jauh panjang dari slice dari awal hingga akhir misal slice[4:6] berarti dia punya length
// 6 - 4 = 2 , 2 merupakan length dari slice tersebut

// capacity yaitu dimana  panjang array- dimana slice dimulai, misal kita punya array dengan panjang 10, kita ambil slice di slice[4:6],
// maka capacitynya akan menjadi 10-4, yatitu 6'

// jika length melebihi capacity maka dia akan menjadi array baru, dan tidak akan bererngaruh ke array yg merepresentasikan slicenya
// kamu bisa mewmbuat slice dengan cara
/*
/	untuk akses semua isi dari slice kamu bisa memakai slice[:]
	untuk mengambil nilai dari index pertama sampai batas yang di tentukan kamu bisa memakai slice[:8], misal sampai index 7
	untuk mengambil nilai dari pointer atau titik awal sampai teralhir kamu bisa memakai slice[2:] misal dari titik awal index 2
	untuk mengambil dari pointer dan length yang di inginkan kamu bisa memakai slice[2:7], ini akan mengambil nilai dari index 3 sampai index 6

/	untuk membuat slice baru kamu bisa memakai format penulisan
	var students = make([]string, length, capacity)

	untuk membuat copy data dari slice kamu bisa memakai
	// jadi kamu haru punya destinationny dulu, atau tempat(wadah )untuk manaruh hasil copy
	copy(destination, sliceCopy)

	menambahkan data baru kedalam slice
	var slice = append(sliceYangSUdahAda, "data baru")

/	untuk mendapatkan panjang dari capacity kamu bisa menggunakan perintah
		cap(slice)

	untuk mendapatkan panjang dari legth kamu bisa mengguanakan perintah
			len(slice)

	/// PENTING jangan salah membedakan array dengan slice
	var name = [...]string{"kuro", "neko", "boku", "ore"}  -> ini adalah array
	var waifu = []string{"kato megumi", "tsukasa chan", "ranko kanzaki", "fuuka miyazawa"}  -> ini adalah slice

*/
// ini contoh implementasinya
func SliceAndFriends() {
	// membuat array nama bulan
	var months = [...]string{"januari", "feruari", "maret", "april", "mei", "juni",
		"juli", "agustus", "september", "november", "oktober", "desember"}

	// mengambil semua nilai dalam slice
	getAllMonths := months[:]
	fmt.Println(getAllMonths)

	// mengambil dari index ke 4 sampai index ke 8
	getFrom4to8 := months[4:9]
	fmt.Println(getFrom4to8)

	// menagmbil panjang capacity dari getFrom4to8, dari 12 - 4 = 8, panjang array - start slice
	getCapacity := cap(getFrom4to8)
	fmt.Println(getCapacity)

	// mengambil panjang slice, dari maxSLice - startSLice, 9 - 4 = 5
	getLength := len(getFrom4to8)
	fmt.Println(getLength)

	// kalau kamu merubah nilai dari slice maka nilai dari array juga akan berubah misal
	// mengambil array dari index ke 10 sampai terakhir
	get10toEnd := months[10:]
	fmt.Println(get10toEnd)
	// nah disini akan coba diubah, oh iya dia akan dimulai dari index ke 0 dari slice
	get10toEnd[0] = "oktober change"
	// disini niali dari index 10  akan berubah, begitupun dengan yang ada di dalam array, dia akan menampilkan index 10  -> oktober change
	fmt.Println(get10toEnd)
	fmt.Println(months)

	// jika misal nilai dari array yang di ubah maka nilai dari slice juga akan berubah misal
	months[0] = "change januari"
	// mengambil nilai dari slice 0-4
	getJanuariChange := months[:5]
	fmt.Println(getJanuariChange)

	// nah bagaimana jika kita tambah tapi melebihi kapasitas dari si array,
	// dia akan membuat slice baru dan tidak akan berpengaruh ke array yang lama, misal
	// mengambil nilai array  10 - end
	tryChange := months[10:]
	tryChangeNow := append(tryChange, "bulan ke 13")

	fmt.Println(tryChangeNow) // disini dia akan membuat slice baru yang tidak akan berpengaruh ke array utama
	fmt.Println(months)       // nah disini dia tidak akan berubah

	// membuat slice baru dengan kapasitas 6 dan length 4
	var mywife = make([]string, 4, 6)
	mywife[0] = "Kaguya Shinomiya"
	mywife[1] = "Ranko Kanzaki"
	mywife[2] = "Chino Kaffui"
	mywife[3] = "Kato Megumi"

	fmt.Println(mywife)

	// nah disini coba kita copy
	// membuat slice baru untuk destination
	var newMyWIfe = make([]string, len(mywife), cap(mywife))
	// copy dari mywife
	copy(newMyWIfe, mywife)
	fmt.Println(newMyWIfe)

	// jangan salah membedakan arra dengan slice
	arr := [...]string{"ko", "mi", "san"}            // ini array
	slc := []string{"ranko", "kato", "megu", "guya"} // ini slice

	fmt.Println(arr, slc)

}
