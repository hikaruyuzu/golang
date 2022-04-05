package gopackage

import (
	"fmt"
	"time"
)

// timeadalah package yang digunakan untuk management waktu di golang
// kamu bisa baca dokumentasinya di https://pkg.go.dev/time#section-documentation

//example
func GetTime() {
	now := time.Now()
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())

	// untuk set  manual, ke tipe date time
	dateNow := time.Date(2002, 8, 10, 10, 10, 10, 10, time.UTC)
	fmt.Println(dateNow)

	// untuk parse string ke date time
	// standarisasi layout konvertion
	layout := time.RFC3339
	parseTime, _ := time.Parse(layout, "2022-02-15T15:04:05+07:00")
	// if err != nil {
	// 	panic(err.Error())
	// } else {
	// }

	fmt.Println(parseTime)
}
