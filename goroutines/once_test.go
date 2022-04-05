package goroutines

import (
	"fmt"
	"sync"
	"testing"
)

// sync once
// hanya akan mengeksekusi program goroutines tepat sekali saja
// jika memiliki kebutuhan mengeksekusi program dan hanya boleh sekali saja maka gunakan sync once

// exmaple
var counter = 0

func RunOnce() {
	counter++
}

func TestOnce(t *testing.T) {
	once := sync.Once{}
	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go func() {
			wg.Add(1)
			// once tidak bisa memanggil func yang memiliki parameter
			// gnakan Do untuk menggunakan once
			once.Do(RunOnce) // your function here
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Done...", counter)
}
