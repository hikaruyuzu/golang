package goroutines

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

// atomic ini digunakan untuk menangani masalah manipulasi data primitife seperi number
// sebelumnya kita menggunakan mutex / RWmutex, untuk mengatasi problem race cndition  akan tetapi ada yang lebih baik untuk menanganinya
// jika data yang kita manipulasi adalah data primitife
// atomic ada di package atomic

// example
func TestAtomic(t *testing.T) {
	var counterAtomic int64 = 0
	var group = sync.WaitGroup{}

	for i := 1; i <= 1000; i++ {
		go func() {
			group.Add(1)
			for j := 1; j <= 100; j++ {
				atomic.AddInt64(&counterAtomic, 1)
			}
			group.Done()

		}()

	}
	group.Wait()

	fmt.Println("Counter ", counterAtomic)
}
