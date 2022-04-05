package goroutines

import (
	"fmt"
	"sync"
	"testing"
)

// sync map sama hanya dengan map map pada umumnya akan teatapi yang membedakannya adalah dia di khususkan untuk cuncoruncy(goroutine)
// didalam map kita bisa memasukan data key dan value seperi halnya pada map biasa
// Store(key, value) -> memasukan data
// Load(key) -> mengambil data
// Delete(key) -> Menghapus data
// Range(func(key, value)interface{}) -> untuk iterasi data, eturn berupa bolean

// example
func AddDataSyncMap(data *sync.Map, value int, wg *sync.WaitGroup) {
	defer wg.Done()
	wg.Add(1)
	// add data
	data.Store(value, value)
}

func TestAddDataSyncMap(t *testing.T) {
	// sync map
	data := &sync.Map{}
	// wait group
	waitGroup := &sync.WaitGroup{}

	// iterate 100 goroutines
	for i := 0; i < 100; i++ {
		go AddDataSyncMap(data, i, waitGroup)
	}
	waitGroup.Wait()

	fmt.Println(data.Load(2))

	data.Range(func(key, value interface{}) bool {
		fmt.Println(key, "=", value)
		return true
	})
}
