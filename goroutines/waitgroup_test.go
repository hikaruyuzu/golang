package goroutines

import (
	"fmt"
	"sync"
	"testing"
)

// sync WaitGroup
// kamu bisa menggunakan sync.WaitGrub untuk menggantikan time.sleep,
// jika kamu punya beberapa goroutine yang harus di eksekusi dan gamau aplikasinya selesai dulu sebelum goroutinenya di eksekusi semua
// kamu bisa menggunakan waigrub ini

// add digunakan untuk menentukan berapa banyak goroutine yang akan di eksekusi
// selalu gunakan done unruk mengurangi nilai dari add
// wait akan menunggu sampai add bernilai 0, dan done bernilai jumlah dari add
// kenapa ?, karena add akan selalu di setaip goroutine di eksekusi dan akan di masukan ke done
// maka dari itu kita harus selalu done setaip membuat waitgrub agar tidak tejadi deadlock

// example
func RunAsAsycronous(group *sync.WaitGroup) {
	// always done wg
	defer group.Done() // akan deadlock jiga ga di done kan
	// amount goroutines
	group.Add(1)
	// line code
	fmt.Println("Waiting ...")
}

func TestWaitGroup(t *testing.T) {
	// create wg
	counter := 0
	var wGroup = sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		go RunAsAsycronous(&wGroup)
		counter++ // cuma buat cek apakah benar"selesai semua
	}
	// nah disini kamu bisa gunakan wait dari waitgroup alih" menggunakan time.sleep
	wGroup.Wait()
	fmt.Println("Done execute all goroutines", counter)
}
