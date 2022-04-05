package helper

import "fmt"

// closuure adalah blok code yang akan di eksekusi/ yang dapat di akses dalam sebuah block kode progarm

// misal
func ClosureTest() {
	var name string
	func() {
		name = "kaguya" // nah kamu bisa mengakses dan mengubah nama dari sini

		// sedangkan jika kamu mencoba membuat kode disini, kamu ga akan bisa mengaksesnya dari luar block {}
		title := "wakaranai"
		// bisanya hanya di akses disini dan di dalam function ini saja
		fmt.Println(title)
	}()

	// fmt.Println(title) -> ini gabisa
	fmt.Println(name)
}
