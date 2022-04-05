package helper

import (
	"fmt"
	"log"
	"testing"
)

// kamu bisa mneggunakan benchmark untuk mengukur kecepatan program
// N(banyak perulangan) akan otomatis ditentukan golang
// golang akan secara otomatis menentukan berapa banyak perulangan yang harus dilakukan dan menentukan hasil kecepatan kode perogram kita

/**
/ gunakan "go test -v -run=NamaTestYangGada -bench=. -> untuk menjalankan semua benchmark
/		  "go test -v -run=NamaTestYangGada -bench=NamaSpesifikBenchmark -> untuk menjalankan benchmark secaa spesifik
		  "go test -v -run=NamaTestYangGada -bench=NamaSpesifikBenchmark/SubName -> akses sub benchmark

*/

// becnhmark juga mendukung table benchmark
func BenchmarkTable(b *testing.B) {
	test := []struct {
		name     string
		request  string
		expected bool
	}{
		{
			name:    "neko1",
			request: "katak",
		},
		{
			name:    "neko2",
			request: "neko",
		},
		{
			name:    "neko3",
			request: "kuro",
		},
		{
			name:    "neko4",
			request: "allo",
		},
		{
			name:    "neko5",
			request: "owwwwwo",
		},
		{
			name:    "neko6",
			request: "woooghhhhseeee",
		},
	}
	// looping semua data dalam struct
	for _, benchmark := range test {
		b.Run(benchmark.name, func(b *testing.B) {
			var result bool
			for i := 0; i < b.N; i++ {
				result = Palindrome(benchmark.request)
			}
			log.Println(result)

		})
	}
}

// benchmark mendukung sub benchmark, yaitu melakukan benchmark dalam satu function yang sama
func BenchmarkPalindromeSub(b *testing.B) {
	b.Run("KAWAI", func(b *testing.B) {
		// code benchamrk
		var result bool
		for i := 0; i < b.N; i++ {
			result = Palindrome("KAWAI")
		}
		log.Println(result)
	})
	// sub
	b.Run("OWO", func(b *testing.B) {
		// code benchamrk
		var result bool
		for i := 0; i < b.N; i++ {
			result = Palindrome("OWO")
		}
		log.Println(result)
	})

}

// awali functionnya dengan nama Benchmark
func BenchmarkPalindrome(b *testing.B) {
	var result bool
	// cukup lakukan looping sampai b.N
	for i := 0; i < b.N; i++ {
		result = Palindrome("KATAK")
	}
	fmt.Println(result)
}

func BenchmarkPalindromeFail(b *testing.B) {
	var result bool
	// cukup lakukan looping sampai b.N
	for i := 0; i < b.N; i++ {
		result = Palindrome("ONICHAN")
	}
	fmt.Println(result)
}
