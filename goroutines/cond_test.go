package goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// cond atau condition ini sama halnya dengan mutex / RWmutex
// untuk membuat cond kita juga harus menggukana locker seperti mutex/rw mutex
// yang membedakan disini dia mempunyai func wait yang digunakan untuk menunggu konvirmasi setelah locking agar prosesnya dapat berjalan
// signal -> digunakan untuk mengirim sinyal(konfirmasi) ke wait untuk melanjutkan eksekusi setlah locking tapi 1 per 1
// jadi misal ada 10 goroutines dia akan di cek 1 per 1

// broadcast -> berbeda dengan signal, disini dia hanya melakukan konvirmasi di awal saja, setelah itu semua program goroutines akan di eksekusi

// esample
var mutex = sync.Mutex{}
var cond = sync.NewCond(&mutex)

var group = sync.WaitGroup{}

func WaitCondition(value int) {
	defer group.Done()
	group.Add(1)

	cond.L.Lock() // locking
	cond.Wait()   // nah disini dia akan meununggu sampai ada konvirmasi signal ataupun broadcast yang masuk ke cond
	// your code here
	fmt.Println("Condition ", value)

	cond.L.Unlock() // unlock
}

func TestCondition(t *testing.T) {
	for i := 1; i <= 10; i++ {
		go WaitCondition(i)
	}

	for i := 1; i <= 10; i++ {
		func() {
			// jadi disini kamu bisa masukin logic apa saja sebelum konformasi eksekusi goroutines
			time.Sleep(1 * time.Second)
			cond.Signal() // mengirim konfirmasi ke WaitCondition (wait) untuk melanjutkan goroutines yang di lock
		}()

	}

	group.Wait()
	// func() {
	// 	// jadi disini kamu bisa masukin logic apa saja sebelum konformasi eksekusi goroutines
	// 	time.Sleep(1 * time.Second)
	// 	cond.Broadcast() // kalau ini dia akan di cek dan melakukan kknfirmasi di awal  ekssekusi goroutines saja, seterusnya tidak
	// }()
}
