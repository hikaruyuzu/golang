package gopackage

import (
	"container/ring"
	"fmt"
	"strconv"
)

// package kontainer ring adalah implementasi structur data ring
// kamu bisa lihat selengkapnya di sini https://pkg.go.dev/container/ring

// contoh penggunaan
func GolangRing() {
	// 10 adalah jumlah datanya
	var data = ring.New(10)
	// interare untuk push data
	for i := 0; i < data.Len(); i++ {
		data.Value = "Data ring " + strconv.FormatInt(int64(i), 10)
		data = data.Next()
	}

	// untuk akses
	data.Do(func(value interface{}) {
		fmt.Println(value)
	})
}
