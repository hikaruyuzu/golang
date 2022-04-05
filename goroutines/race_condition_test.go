package goroutines

import (
	"fmt"
	"testing"
	"time"
)

// race condition, adalah dimana jika kita memanipulasi data yang sama dalam satu waktu, dalam beberapa(banyak process goroutines)
// ini akan menyebabkan problem dengan nama race condition, dimana satu variable akan di shareing dalambeberapa goroutines dalam
// waktu yang bersamaan, hal ini menjadi masalah dalam cuncrrency karena ada beberapa thread yang akan di eksekusi secara parallel

// example problem race condition
func TestRaceCondition(t *testing.T) {
	var counter = 0
	for i := 1; i <= 1000; i++ {
		// execute with goroutines
		go func() {
			for i := 1; i <= 100; i++ {
				// disini bisa terjadi race condition
				counter = counter + 1
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("counter now ", counter)
}
