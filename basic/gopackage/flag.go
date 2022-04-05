package gopackage

import (
	"flag"
	"fmt"
	"os"
)

// package flag erat kaitannya dengan package os
// biasanya digunakan untuk membuat aplikasi berbasis command line

// kamu bisa baca dokumentasinya disini https://pkg.go.dev/flag

// example
func ExecuteWithCommandLine() {
	// ini dapat di set lewat commandline
	// cara akses, panggil program di main program
	// gunakan perintah go run main.go -host=namahost -password=enterpasswordnya
	
	hostname := flag.String("host", "kato megumi", "put your hostname")
	password := flag.String("password", "katochan", "put your password")

	// harus selalu dipanggil agar ga error
	flag.Parse()

	os.Setenv("HOSTNAME", *hostname)
	os.Setenv("PASSWORD", *password)

	fmt.Println("HOSTNAME KAMU: ", os.Getenv("HOSTNAME"))
	fmt.Println("PASSWORD KAMU: ", os.Getenv("PASSWORD"))

}
