package main

import (
	"fmt"
	"time"
)

func main() {
	orders := make(chan string, 3) // Buffered channel with a capacity of 3

	// Finder placing mining location
	go func() {
		for i := 1; i <= 5; i++ {
			order := fmt.Sprintf("Mining order #%d", i)
			orders <- order // Blocks until the miner is ready to accept new order
			fmt.Println("🗒️ Placed:", order)
		}
		close(orders)
	}()

	// Miner processing orders
	for order := range orders {
		fmt.Printf("🔨 Mining: %s\n", order)
		time.Sleep(2 * time.Second) // Time taken to mining
		fmt.Printf("🎉 Done: %s\n", order)
	}
}
