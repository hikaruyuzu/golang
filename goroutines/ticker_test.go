package goroutines

import (
	"fmt"
	"testing"
	"time"
)

// kamu bisa menggunakan ticker untuk melakukan job yang berulang setelah beberapa waktu yang di tentukan, berbeda halnya dengan timer

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second) // akan mengulangi job setiap 1 second

	// untuk mengentikan ticker kamu bisa menggunakan ticker.Stop
	go func() {
		time.Sleep(5 * time.Second)
		ticker.Stop()
	}()

	for time := range ticker.C { // umumnya ini dilakukan dengan switch, agar tidak terjadi dead lock
		fmt.Println("Do something here ", time)

	}

}

func TestTickTicker(t *testing.T) {
	// disini kita tidak usah mengembalikan data channelnya C
	tick := time.Tick(2 * time.Second)

	for tickTick := range tick {
		fmt.Println("tick ", tickTick)
	}
}
