package helper

import "fmt"

// digoalng kita bsa mmbuat function sebagai paramater yang artinya kita bisa memasukan function kedalam sebnuah parameter

// misal
func GetMyName(name string) string {
	return fmt.Sprintf("Hello my name is %s", name)
}

// pada saat pemanggilan kamu bisa memasukan GetMyName sebagai value dari sebuah parameter
/*
/ misal
	YourName := GetMyName
	result := YourName("Kaguya shinomiya") -> nah kamu bisa memakai YourName ini sebagai function
*/