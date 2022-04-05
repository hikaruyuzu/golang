package goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// pool adalah implementasi design pattern yang bernama object pool patern
// biasanya dalam kasus nyata ini digunakan untuk manage koneksi ke database, misal untuk menampung 10 koneksi
// hanya digunakan dalam cuncorency
// tidak akan mengalami race condition
// setelah menggunakan data dalam pool kamu bisa menaruhnya kembali dengan cara menggunakan put
// get digunakan untuk mendapatkan data dari pool
// jika data dalam pool di ambil dengan method get maka data dalam pool akan kosong

func TestCreatePool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} { // kamu bisa membuat nilai defult jika data nil menggunakan New : func anonymus dengan return interface
			return "Ara ara ~ kawaiii"
		},
	}

	// add data ke pool
	pool.Put("Kaguya sinomiya")
	pool.Put("Tsukasa chan")
	pool.Put("Kato Megumi")

	for i := 1; i <= 10; i++ {
		go func() {
			// mengambil data dari pool
			data := pool.Get()

			time.Sleep(1 * time.Second) // gambaran jika di delay, data telat di kembalikan maka pool akan mengembalikan data nill

			fmt.Println(data)

			// menaruh data setelah digunakan
			pool.Put(data)
		}()
	}

	time.Sleep(3 * time.Second)
}
