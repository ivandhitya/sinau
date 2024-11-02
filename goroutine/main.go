package main

import (
	"fmt"
	"time"
)

func berhitung(nama string, sleep int) {
	for i := 0; i < 5; i++ {
		fmt.Printf("%s : %v\n", nama, i)
		time.Sleep(time.Millisecond * time.Duration(sleep))
	}
}
func main() {
	go berhitung("abdi", 200)
	berhitung("ajeng", 400)
}
