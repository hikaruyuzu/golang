package goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// timer dilakukan untuk melakukan delayjob dekali saja,berbeda dengan ticker yang akan terus berulang
// misal ada suatu function yang kita ingin eksekusi % detik setelah function tersebut di eksekusi, kamu bisa menggunakan timer

func TestTimer(t *testing.T) {
	timeTimer := time.NewTimer(5 * time.Second) // menungu, akan di eksekusi 5 detik kemudian
	fmt.Println(time.Now())

	timeT := <-timeTimer.C // akan di eksekusi 5 detik kemudian
	fmt.Println(timeT)
}

func TestTimeAfter(t *testing.T) {
	timeAfter := time.After(5 * time.Second)
	fmt.Println(time.Now())
	after := <-timeAfter
	fmt.Println(after)
}

func TestAfterFunc(t *testing.T) {
	group := sync.WaitGroup{}
	// disini kita bisa langsung passing saja functrionnya kedalam timer
	group.Add(1)
	time.AfterFunc(5*time.Second, func() {
		fmt.Println("Do something here after 5 second") // akan di eksekusi setelah5 second,  akan di eksekusi as goroutines
		group.Done()
	})
	group.Wait()
	fmt.Println(time.Now())

}
