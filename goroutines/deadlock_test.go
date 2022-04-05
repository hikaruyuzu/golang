package goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// deadlock adalah masalah yang sering terjai jika kita menggunakan goroutines

// ini dalah contoh kasus deadlock, dimana goroutines akan saling tunggu satu sama lain

type UserBalance struct {
	Mutex   sync.Mutex
	Name    string
	Balance int
}

// lock
func (userBalance *UserBalance) Lock() {
	userBalance.Mutex.Lock()
}

// unlock
func (userBalance *UserBalance) Unlock() {
	userBalance.Mutex.Unlock()
}

// send amount
func (userBalance *UserBalance) ChangeBalance(amount int) {
	userBalance.Balance = userBalance.Balance + amount
}

// transfer, disini akan terjadi deadlocck jika ada dua user yang melakukan transfer secara bersamaan
func TransferBalance(userSender *UserBalance, userReceiver *UserBalance, amountTransfer int) {
	// lock first
	userSender.Lock()
	userSender.ChangeBalance(-amountTransfer)
	fmt.Println("Lock sender user 1, ", userSender.Name)

	time.Sleep(1 * time.Second)

	userReceiver.Lock()
	fmt.Println("Lock receiver user 2", userReceiver.Name)
	userReceiver.ChangeBalance(amountTransfer)

	// unlock after transfer
	userSender.Unlock()
	userReceiver.Unlock()

}

// nah disini kita akan simulasikan deadlock
func TestDeadLock(t *testing.T) {
	user1 := UserBalance{
		Name:    "Kaguya shinomiya",
		Balance: 2000000,
	}

	user2 := UserBalance{
		Name:    "Ranko Kanzaki",
		Balance: 2000000,
	}

	// disini akakn terjadi deadlock, dimana goroutine 1 dan goroutine 2 akan saling lock dan tidak akan bisa lock user yang lain direceiver
	// karena sudah terlanjur di lock di awal
	// akan saling tunggu, dan terjadi deadlock

	// disini user1 akan di lock, dan tidak bisa lock user2 yang sudah terlanjur di lock di goroutines2
	go TransferBalance(&user1, &user2, 500000)
	// disini user2 akan di lock dan tisak bisa lock user1 yang sudah terlanjur di lock di goroutine 1
	go TransferBalance(&user2, &user1, 200000)

	time.Sleep(5 * time.Second)
	fmt.Println(user1.Name, "Balance now ", user1.Balance)
	fmt.Println(user2.Name, "Balance now ", user2.Balance)

}
