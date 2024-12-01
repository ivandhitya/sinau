package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go berhitung("abdi", 200, &wg)
	wg.Add(1)
	go berhitung("ajeng", 400, &wg)
	wg.Wait()
}

func berhitung(nama string, sleep int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Printf("%s : %v\n", nama, i)
		time.Sleep(time.Millisecond * time.Duration(sleep))
	}
}
