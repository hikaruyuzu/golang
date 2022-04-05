package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

/**
run unit test
	go test -v -run=TestFunctionName

run sub unit test
	go test -v-run=TestFunctionName/SubTest

run all unit test
	go test -v

run all unit test from other package
	go test -v ./...
*/

// for better experience, gunakan table test agar tidak perlu mengulang unit test yang sama
// gunakan konsep perulangan dan sub test untuk membuat sub test
// masukan kode yang ingin di test kedalam struct of slice
func TestPlindromeTableTest(t *testing.T) {
	// untuk menampung expected dan request
	test := []struct {
		name     string
		request  string
		expected bool
	}{
		{
			name:     "neko1",
			request:  "katak",
			expected: true,
		},
		{
			name:     "neko2",
			request:  "neko",
			expected: false,
		},
		{
			name:     "neko3",
			request:  "kuro",
			expected: false,
		},
		{
			name:     "neko4",
			request:  "allo",
			expected: true,
		},
	}
	// gunakan looping
	for _, value := range test {
		t.Run(value.name, func(t *testing.T) {
			result := Palindrome(value.request)
			require.Equal(t, value.expected, result, "some messages")
		})
	}
}

// sub unit test, membuat sub dudalam function test bisa lebih dari 1
// gunakan t.Run("nameFunction", func anonymus)
func TestPalindromeSub(t *testing.T) {
	t.Run("NEKO", func(t *testing.T) {
		// some testing
		result := Palindrome("neko")
		assert.Equal(t, false, result, "some messages")
		fmt.Println("done...")
	})

	t.Run("YUKUY", func(t *testing.T) {
		// some testing 2
		result := Palindrome("yukuy")
		assert.Equal(t, true, result, "some messages")
		fmt.Println("done...")
	})
}

// before and after test using testing.M
// akan di eksekusi per 1 package
// before dan after akan di eksekusi sebelum dan sesudah ubit test /package
// bisa di isi logic apapun
// wajib menggunakan nama func TestMain

func TestMain(m *testing.M) {
	// do something before running unit test
	fmt.Println("DO SOMETHING BEFORE UNIT TEST")
	// use m.Run() to execute all program unit test
	m.Run() // -> semua program di eksekusi
	// do something afer running unit test
	fmt.Println("DO SOMETHING AFTER UNIT TEST")

}

// skip test, digunakan untuk mengkip unit test di suatu kondisi misal kode hanya jalan di windows, etc
func TestPalindromeSkip(t *testing.T) {
	result := Palindrome("moka")
	fmt.Println(runtime.GOOS)
	if runtime.GOOS != "darwin" {
		t.Skip()
	}
	// never execute
	assert.Equal(t, false, result, "some messages")
	fmt.Println("done execute")
}

// assert, dari package testify
func TestPalindromeWithAssert(t *testing.T) {
	result := Palindrome("moka")
	// parameter, false is expect
	assert.Equal(t, false, result, "some message")
	// after assert, akan di eksekusi, like fail
	fmt.Println("will execute")
}

// require , dari package testify
func TestPalindromeWithRequire(t *testing.T) {
	result := Palindrome("moka")
	require.Equal(t, true, result, "some message")
	// after assert, program tidak akan di eksekusi, like fail now
	fmt.Println("never execute")
}

// unit test palindome
// fail, jika terjadi gagal (error) pada kode progra kode akan di lanjutkan ke kode berikutnya
// failNow , jika terjadi gagal (error) pada kode perogram kode akan di berhentikan saat itu juga
// error, berkaitan erat dengan fail, akan tetapi dia akan bisa di beri error exception, dan setelah itu dia akan mengeksekusi fail
// fatal , berkaitan erat dengan failNow sama seperti error bedanya dia bisa dikasih error exception, lalu mengeksekusi fail now
func TestPalindrome(t *testing.T) {
	result := Palindrome("katak")
	if result != true {
		// kamu bisa menambahkan error message jika menggunakan error/ fatal
		t.Error("result must be 'true'")
	}
	// akan di eksekusi jika menggunakan fail/ error
	// like fail
	fmt.Println("test done")
}

// > 1 unit test
func TestPalindromeVailed(t *testing.T) {
	result := Palindrome("kata")
	if result != false {
		// like fail now
		t.Fatal("result must be 'false'")
	}
	// akan di ignore
	fmt.Println("test done")
}
