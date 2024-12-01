package main

import (
	"fmt"
	"sync"
)

var counter int = 0
var mutex = &sync.Mutex{}

func increment(wg *sync.WaitGroup) {
	mutex.Lock() // Mengunci akses ke counter
	counter++
	mutex.Unlock() // Membuka akses setelah selesai
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment(&wg)
	}

	wg.Wait()
	fmt.Println("Final Counter:", counter)
}
