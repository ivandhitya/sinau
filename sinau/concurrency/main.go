package main

import (
	"fmt"
	"runtime"
	"time"
)

func process(name string, duration time.Duration) {
	fmt.Println("#Process " + name + " started")
	time.Sleep(duration)
	fmt.Println("#Process " + name + " done")
}

func main() {
	// program hanya berjalan dengan menggunakan 1 processor
	runtime.GOMAXPROCS(1)

	// tugas pertama
	go func() {
		process("1A", 2*time.Second)
		process("1B", 1*time.Second)
	}()

	// tugas kedua
	go func() {
		process("2A", 1*time.Second)
		process("2B", 3*time.Second)
	}()

	time.Sleep(5 * time.Second)
}
