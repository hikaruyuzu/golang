package gopackage

import (
	"fmt"
	"os"
)

// package os digunakan untuk apa apa saja yang berhubungan dengan sistem operasi
// baca dokumentasi di https://pkg.go.dev/os

// contoh penggunaan
func GetArguments() []string {
	data := os.Args
	return data

	// cara akses
	// data := gopackage.GetArguments()
	// fmt.Println(data[1])
	// fmt.Println(data[2])

	// fmt.Println(data)
}

// contoh lain set environment
func SetEnv(username string, password string) {
	os.Setenv("USERNAME", username)
	os.Setenv("PASSWORD", password)
}

// contoh get environment
func GetEnv() {
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	fmt.Println(username)
	fmt.Println(password)

	/**
	gopackage.SetEnv("Kato megumi", "root")
	gopackage.GetEnv()
	*/

}

// untuk emndapat hostname
func GetHostname() {
	host, err := os.Hostname()
	if err != nil {
		fmt.Println("ERROR ", err.Error())
	} else {
		fmt.Println(host)
	}
}

// masih banyak bangettt baca dokumentasii
