package helper

import "fmt"

// kamu juga bisa membuat funcrtin secara langsung tanpa harus membuat functionny satu persatu ini dinamaakn function annoymus
// contoh

// alias parameter
type BlackList func(string) bool

func BlockUser(name string, block BlackList) {
	if block(name) {
		fmt.Println("Kamu di block ", name)
	} else {
		fmt.Println("Hello", name)

	}
}

// kamu tidak perlu membuat pengecekan satu persatu untuk function disini kamu bisa membuatnya secara langsung ketika pemanggilan function
/*
	/ contoh emanggilan function
	blaclist := func(name string) bool {
		if name != "admin" {
			return false
		} else {
			return true
		}
	}

	helper.BlockUser("admin", blaclist)

	// kamu juga bisa membuatnya langsung seperti ini
	helper.BlockUser("Ranko kanzaki", func(name string) bool {
		if name == "wakaranai" {
			return true
		} else {
			return false
		}
	})
*/

// adapaun function yang akan selalu di esekusi jika rogram dijaan=nkan misal
func GetVtuber(name string) {
	func() {
		fmt.Println("Hello ", name)
	}()
}
