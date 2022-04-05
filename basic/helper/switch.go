package helper

import (
	"fmt"
	"strconv"
)

// swith hampir sama dengan if akan tetapi dia hanya untuk membandingkan satu variable saja, hanya digunakan untuk operasi yang sederhana
// kamu bisa mendeklarasikan parameter dari swirh secara lsngsung seperti halnya if
// kamu juga bisa melakukan perbandingan serti pada if expression akan tetapi lebih simple, tidak sekompleks di if expression

// contoh function kalkukator sedehana
func SimpleCalculator(operasi string, value1 int, value2 int) string {

	var result int
	switch operasi {
	case "+":
		result = value1 + value2
	case "-":
		result = value1 - value2
	case "*":
		result = value1 * value2
	case "/":
		result = value1 / value2
	default:
		fmt.Println("maaf operasi", operasi, "tidak tersedia")
	}

	resultInString := strconv.Itoa(result)
	return fmt.Sprintf("hasil dari operasi %s %s %s adalah %s", strconv.Itoa(value1), operasi, strconv.Itoa(value2), resultInString)
}

// kamu juga bisa membuat simple operasi perbandingan dengan switch ini, lebih mirip dengan if akan tetapi lebih simple saja
func SImpleComparation(value1, value2 int) string {
	var greetings string
	// disini ga usah memakai parameter
	switch {
	// disini kita bisa memakai operasi perbaingan sederhana, bisa lebih dari satu kondisi namun tidak se kompleks if
	case value1 > value2:
		greetings = fmt.Sprintf("nilai dari %s lebih besar dari %s", strconv.Itoa(value1), strconv.Itoa(value2))

	case value1 < value2:
		greetings = fmt.Sprintf("nilai dari %s lebih kecil dari %s", strconv.Itoa(value1), strconv.Itoa(value2))

	}
	return greetings
}

// kamu juga bisa mendeklarasikan parameter untuk swith secara lsngsung seperti pada if
func GetLengthKataWithSwitch(kata string) string {
	var result string
	// nah disini kamu bisa mendeklarasikan parameter secara langsung
	switch length := len(kata); length {
	case 5:
		result = "panjang kata adalah 5"
	case 10:
		result = "panjang kata dalah 10"
	default:
		result = "panjang kata tidak diketahui"
	}

	return result
}
