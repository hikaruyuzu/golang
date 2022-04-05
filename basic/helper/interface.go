package helper

import (
	"fmt"
	"strconv"
)

// interface adalah tipe data kontrak, jadi setiap kita membuat interface yang berisi function kita harus mndefinisikan/ membuat function
// tersebut agar interface dapat digunakan, jika tidak maka program akan error

// misal interface untuk convert nilai mata uang berikut

// interface akan memanggil function" yang telah dibuat
type ConverttMoney interface {
	// representasi function yang harus dibuat
	ChangeCurrencyValueToIdr(money int) int
}
type Rupiah struct {
	Idr int
}

// memanggil kontrak intergace dalam function dan menggunakan struct method rupiah
// dan ini sebagai main function
func (rupiah Rupiah) ResultConvertInIdr(money ConverttMoney) string {
	result := money.ChangeCurrencyValueToIdr(rupiah.Idr)
	return fmt.Sprintf("Mata uang ke convert %s", strconv.Itoa(result))
}

// struct kontrak from yen
type Yen struct {
	CurrentYen float32
}

// melengkapi kontrak interface, dan implementasi struct YEN
func (yen Yen) ChangeCurrencyValueToIdr(money int) int {
	result := money / int(yen.CurrentYen)
	return result
}

// struct kontrak from dollar
type Dollar struct {
	CurrentDollar float32
}

// melengkapi kontar interface, dan implementasi struct DOLLAR
func (dollar Dollar) ChangeCurrencyValueToIdr(money int) int {
	result := money / int(dollar.CurrentDollar)
	return result
}

// untuk cara pemanggilan kamu harus membuat dulu struct untuk idr dan nilai mata uang untuk tiap struct
/**
/ untuk pemanggialan

 	yen := helper.Yen{
		CurrentYen: 124.89,
	}

	dollar := helper.Dollar{
		CurrentDollar: 14325.55,
	}

	idr := helper.Rupiah{
		Idr: 100000,
	}

	data := idr.ResultConvertInIdr(dollar)
	fmt.Println(data)

*/
