package context

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

// sama saja dengan context cancel ataupun timwout, disini context deadline lebih ke spesifik waktu

func GoroutineContextDeadline(ctx context.Context, wg *sync.WaitGroup) chan int {
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
				time.Sleep(1 * time.Second) // simulasi slow/ lemot
			}
		}
	}()
	wg.Done()

	return forever
}

func TestGoroutineContextDeadline(t *testing.T) {
	// first create context
	ctxbg := context.Background()
	ctx, cancel := context.WithDeadline(ctxbg, time.Now().Add(10*time.Second)) // bats waktu 10 detik

	// alwayas defer close
	defer cancel()

	// waitgroup
	wg := sync.WaitGroup{}
	fmt.Println("Goroutine run ", runtime.NumGoroutine())

	resource := GoroutineContextDeadline(ctx, &wg)
	fmt.Println("Goroutine run ", runtime.NumGoroutine())

	for source := range resource {
		fmt.Println(source)
	}

	wg.Wait()

	time.Sleep(2 * time.Second)

	fmt.Println("Goroutine run ", runtime.NumGoroutine())

}
