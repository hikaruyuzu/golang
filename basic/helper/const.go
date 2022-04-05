package helper

// const ini adalah nilai yang tidak aken bisa di ubah nilainya, jika dipaksa di ubah maka akan menampilkan error
// walau const tidak dipakai dia tidak akan errot
// cara pembuatan const

/*
const name string = "your name here"
const name = "your name"

/ atau bisa menggunakan multiple
const (
	firstName = "kaguya shinomiya"
	age int = 17
)
*/

func CreateFixWaifu() string {
	const (
		name   string = "Kaguya shinomiya"
		age           = 17
		gender        = "female"
		anime         = "Kaguya sama love is war"

		//anime = "gochuman" -> akan error
	)
	// walau yang lain tidak dipakai dia tidak akan error
	return name
}
