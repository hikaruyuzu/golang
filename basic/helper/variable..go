package helper

import "fmt"

// cara mendeklarasikan variable ada banyak caranya, misal
/*
var name string
var name = "enter your name here"
name := "enter your new namw"

/ atau dapat dipersingkat menjadi
var (
	title = "kaguya sama love is war"
	releaseDate = 24-08-2017
)
*/

// fungsi untuk menganti nama default
func ChangeName(name string) string {
	var result string
	if name != "" {
		result = name
	} else {
		result = ""
	}

	return fmt.Sprintf("Hello my name is Koichi, yorosiku ne %s chan", result)
}
