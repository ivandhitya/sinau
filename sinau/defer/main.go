package main

import "fmt"

func Jumlah(x int, y int) (b int) {
	defer func() {
		b++
		fmt.Println(b)
	}()
	b = x + y
	return b
}

func main() {
	Jumlah(1, 1)
}
