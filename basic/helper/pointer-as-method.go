package helper

// kamu bisa menggunakan pointer untuk method function, karena dia akan pass by value
// maka jika mengguanakan pointer itu akan menjadikannya pass by reference
// sangat disaranka menggunakan pointer karena akan lebih hemat memrory

// untuk cara pemanggilan aga berbeda dengan jika pointer di parameter disini kita tidak membutuhkan key, &

// contoh
type MarriedWith struct {
	Name string
}

// contoh jika tanpa pointer,
// disini datanya tidak akan berubah
func (marriedWith MarriedWith) WaifumuWithoutPointer(name string) {
	// mengganti isi struct name, tapi disini tidak akan berubah karena parameterny bukan pointer
	marriedWith.Name = marriedWith.Name + "Married with" + name
}

// ini adalah contoh jika menggunakan pointer
// data yang di ubah akan mengubah data utama
func (marriedWith *MarriedWith) WaifumuWIthPointer(name string) {
	// disini nilai dari isi struct name akan berubah, karena dia merupakan pointer
	marriedWith.Name = marriedWith.Name + "Married with " + name
}

/**
/	cara pemanggilan

	person := helper.MarriedWith{
		Name: "Koichi Ren",
	}

	// disini ga akan berubah karena bukan pointer

	person.WaifumuWithoutPointer("Kaguya Shinomiya")
	fmt.Println(person.Name)

	// disini akan berubah karena menggunakan pointer di method functionnya
	// untuk pemanggilannya kamu tidak usah memakai tand a& lagi, karena dia otomatis sudah tau kalau dia adalah pointer

	person.WaifumuWIthPointer("Kato Megumi")
	fmt.Println(person.Name)
*/
