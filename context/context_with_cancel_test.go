package context

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

// context cancel digunakna untuk menangani connection leak
// connection leak adalah kondisi dimana goroutine terus terusan mengirim data kedalam channel tanpa henti
// dimana kalaupun dia sudah di hentikan dengan break/ return dia akan tetap mengeksekusi perintahnya di func goroutinenya tanpa henti
// ini di karenakan tidak ada yang mengirim sinyal cancel ke dia, kalau kondisinya sudah berhenti
// nah disini kita bisa mneggunakan cancel

// example connection leak
func ContextGoroutineLeak() chan int {
	forever := make(chan int)

	go func() {
		counter := 1
		defer close(forever)
		for {
			forever <- counter
			counter++
		}
	}()

	return forever
}

func TestContextGoroutineLeak(t *testing.T) {

	fmt.Println("Goroutine run ", runtime.NumGoroutine())
	resultLeak := ContextGoroutineLeak()

	for res := range resultLeak {
		fmt.Println("data ", res)

		if res == 10 {
			break
		}
	}

	fmt.Println("Goroutine run ", runtime.NumGoroutine())
	// disini goroutines akan tetap berjalan karena func ContextGoroutinesLeak, akan terus meneruskan processnya
	// disini akan sangat berbahaya jka kita punya rquest banyak / secondnya

}

// nah kamu bisa solve masalah itu memakai context with cancel

func SolveProblemGoroutineLeak(ctx context.Context, wg *sync.WaitGroup) chan int {
	forever := make(chan int)

	wg.Add(1)
	go func() {
		counter := 0
		defer close(forever)

		for {
			select {
			case <-ctx.Done():
				return
			case forever <- counter:
				counter++
			}
		}
	}()
	wg.Done()

	return forever
}

func TestSolveProblemGoroutineLeakWithContextCancel(t *testing.T) {
	// first create context
	ctxbg := context.Background()
	ctx, cancel := context.WithCancel(ctxbg)

	// waitgroup
	wg := sync.WaitGroup{}
	fmt.Println("Goroutine run ", runtime.NumGoroutine())

	resource := SolveProblemGoroutineLeak(ctx, &wg)
	fmt.Println("Goroutine run ", runtime.NumGoroutine())

	for source := range resource {
		fmt.Println(source)

		if source == 10 {
			break
		}
	}

	wg.Wait()
	cancel()

	time.Sleep(2 * time.Second)

	fmt.Println("Goroutine run ", runtime.NumGoroutine())

}
