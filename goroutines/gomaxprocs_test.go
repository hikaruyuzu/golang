package goroutines

import (
	"fmt"
	"runtime"
	"testing"
)

// bisa digunakan untuk melihat ataupun mengubah jumlah thread
// by default banyak thread akan sama dengan banyak thread di cpu kita

func TestGomaxprocs(t *testing.T) {

	// mendapatkan jumlah cpu
	amountCpu := runtime.NumCPU()
	fmt.Println("cpu ", amountCpu)

	// runtime.GOMAXPROCS(20)                 // disini akan mengubah jumlah thread tapi ini sangat jarang dilakukan
	amountThread := runtime.GOMAXPROCS(-1) // jika kita ambil > 0 itu artinya dia akan mengubah jumlah thread
	fmt.Println("thread ", amountThread)

	numberOfRunningGoroutines := runtime.NumGoroutine() // by default akan berjalan 2 untuk running program ini dan manage garbage collector
	fmt.Println("Goroutine running", numberOfRunningGoroutines)
}
