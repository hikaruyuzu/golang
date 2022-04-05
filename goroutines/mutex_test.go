package goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// sync Mutex
// untuk menyelesaikan masalah race condition kamu bisa menggunakan sync.mutex
// dimana ini dapat digunakan jika kita melakukan shareing variable dalam beberapa goroutines
// kita bisa menggunakan lock dan unlock

// menyelesaikan masalah race condition dengan mutex
func TestMutexRaceCondition(t *testing.T) {
	var counter = 0
	var mutex sync.Mutex
	for i := 1; i <= 1000; i++ {
		// execute with goroutines
		go func() {
			for i := 1; i <= 100; i++ {
				// disini agar tidak terjadi race condition kita bisa lakukan lock dan unlock dengan mutex
				mutex.Lock()
				counter = counter + 1
				mutex.Unlock() // untuk melepas lock dan melanjutkan lock ke goroutine selanjutnya
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("counter now ", counter)
}

// RWmutex sebenarnya sama saja dengan mutex biasa akan tetapi disini kita bisa melakukan lock read dan write denga satu RWmutex
// sebenarnya dengan menggunakan mutex biasa saja sudah bisa, akan tetapi akan membutuhkan 2 mutex untuk melakukan locking dan unlock

// example RWmutex
type BankAccount struct {
	RWmutex sync.RWMutex
	Balance int
}

func (bankAccount *BankAccount) AddAmount(amount int) {
	// RWmutex lock and unlock, untuk write
	bankAccount.RWmutex.Lock()
	bankAccount.Balance = bankAccount.Balance + amount
	bankAccount.RWmutex.Unlock()
}

func (bankAccount *BankAccount) GetBalance() int {
	// RWmutex rlock and runlock, untuk membaca
	bankAccount.RWmutex.RLock()
	balance := bankAccount.Balance
	bankAccount.RWmutex.RUnlock()

	return balance
}

func TestRwMutex(t *testing.T) {
	// struct bank account
	var backAccount = BankAccount{}

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				// akan aman dari race condition
				backAccount.AddAmount(1)
				fmt.Println(backAccount.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Total balance ", backAccount.GetBalance())

}
