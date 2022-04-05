package gopackage

import "fmt"

// dalam satu pakage tidak bolehada function dengan nama yang sama
// jika ada function dengan nama yang sama maka akan terjadi error
// nama package umumnya akan sama dengan nama folder
// artinya package = folder
// kamu bisa mengakses function dalamsebuah package selagi function tersebut adalah function Global, dengan ciri di awali dengan
// huruf besae

// gunakan keyword import "nama-folder-project/nama-package"   -> untuk mengakses
// gunakan nama-package.NamaFunction ntuk menggunakan function

// contoh
func GetName(name string) string {
	return fmt.Sprintf("Hello %s", name)
}

/**
// wajib melakukan import
import (
	"basic/gopackage"
	"fmt"
)

*/

/**
	/ berikut adalah cara aksesnya

	import (
		"basic/gopackage"
		"fmt"
	)

func main() {

	data := gopackage.GetName("Kaguya shinomiya")
	fmt.Println(data)

}


*/
