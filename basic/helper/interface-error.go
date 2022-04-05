package helper

import "errors"

// error merupakan kondisi pengembalian data bersama data utama jika di dalam progarm terjadi error
// error merupakan representasi dai interface

// contoh program

func DivideErrors(number int, divided int) (int, error) { // disini mengembalikan 2 tipe data yaitu int dan error yan berupa interface
	if divided == 0 {
		return 0, errors.New("Maaf pembagi tidak boleh 0")
	} else {
		var result = number / divided
		return result, nil
	}

}

// cara pemanggilan
/**
data, err := helper.DivideErrors(20, 0)
	if err != nil {
		fmt.Println("Error ", err.Error())
	} else {
		fmt.Println(data)
	}
*/
