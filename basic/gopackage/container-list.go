package gopackage

import (
	"container/list"
	"fmt"
)

// merupakan implementasi struktur data double linkedlist
// kamu bisa baca dokumentasinya di https://pkg.go.dev/container/list

// contoh
func GolangDoubleLinkedlist() {
	data := list.New()
	// digunakan untuk menambah data di belakang
	data.PushBack("Kaguya shinomiya")
	data.PushBack("Kato megumi")
	data.PushBack("Tsukasa Chan")
	// untuk push data depan
	data.PushFront("Ranko Kanzaki")

	// iterate data
	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

}
