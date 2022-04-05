package helper

import "fmt"

// hello world dengan golang
func Konichiwa(name string) {
	fmt.Printf("Konichiwa %s \n", name)
}

// jika menggunaka return
func KonichiwaIntroduce(name string, from string) string {
	return fmt.Sprintf("Hello %s from %s my name is koichi", name, from)
}
