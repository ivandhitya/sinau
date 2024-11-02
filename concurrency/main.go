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
	runtime.GOMAXPROCS(1)

	go func() {

		process("1A", 2*time.Second)
		process("1B", 1*time.Second)
	}()

	go func() {
		process("2A", 1*time.Second)
		process("2B", 3*time.Second)
	}()

	time.Sleep(5 * time.Second)
}
