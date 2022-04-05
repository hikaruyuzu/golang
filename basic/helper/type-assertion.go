package helper

// type asssertion adalah metode mengubah data interface kosong pada return value atau inputan
// sesuai dengan yang di kehenaki
// gunakan keyword result.(nameDataType)

// contoh
func RandomTypeAsesertion() interface{} { // return interface kosong
	return true // disini misal return aluenye benar"random
}

/**
/ kamu bisa mengguanakan type assertion secara langsung jika percaya bahwa tipe datanya adalah misal fix string
/ kamu bisa konversi secara langsung jikan yakin, akan tetapi jika salah konversi maka akan terjadi panic dan akan mengehntikan aplikasi
/ untuk msalah itu kamu bisa menggunakan switch statement denga keyword result.(type) -> untuk mengecek tipe data

// ini contoh yang tidak aman
data := helper.RandomTypeAsesertion() // interface
dataString := data.(string)           // interface konversi to string
fmt.Println(dataString)

// versi aman mengguankan swith case

	result := helper.RandomTypeAsesertion()
	switch data := result.(type) {
	case string:
		fmt.Println("Data string, ", data)
	case int:
		fmt.Println("Data integer, ", data)
	default:
		fmt.Println("wakaranai desu")
	}

*/
