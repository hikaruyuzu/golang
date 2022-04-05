package helper

// struct method akan terlihat sama seperti si strcut memiliki function padahal bukan
// kamu bisa memasukan data field dari struct kedalam function dengan mengakses nama dari stuct.NameOfFunction

// contoh untuk funcstio dan struct untuk menghitung luas lingkaran
// rumus l = pi x r^2

type Circle struct {
	Pi float32
}

// function kamu bisa membuat functionny seperi ini
func (circle Circle) GetLargeCircle(r float32) float32 {
	var result = circle.Pi * r * r
	return result
}

/**
/ kamu bisa memanggil dan menggunakannya dengan cara seperti berikut

circle := helper.Circle{
	Pi: 3.14,
}

data := circle.GetLargeCircle(20192)
fmt.Println(data)

*/
