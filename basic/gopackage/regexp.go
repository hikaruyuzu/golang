package gopackage

import (
	"fmt"
	"regexp"
)

// digunakan untuk valiasi, manipulasi string etc
// baca di https://github.com/google/re2/wiki/Syntax and https://pkg.go.dev/regexp

func CreateRegexp() {
	// membuat regexp baru
	var regex = regexp.MustCompile("a([a-z])o")

	fmt.Println(regex.MatchString("KAto"))
	fmt.Println(regex.MatchString("ako"))
	fmt.Println(regex.MatchString("ato"))
	fmt.Println(regex.MatchString("kuro"))

	// get count true, -1 disini adalah maximal jumlahnya, -1 artinya tidak terbatas, kalau mau membatasi gunakasn angka biasa misal 3, etc
	myRegex := regex.FindAllString("ako kuro kato cino iko alo ego ato tri aso", -1)
	fmt.Println(myRegex)

}
