package main

import "fmt"

func main() {
	ch := make(chan string)
	go sendMessage(ch)

	msg := <-ch
	fmt.Println(msg)
	// atau
	//fmt.Println(<-ch)
}

func sendMessage(ch chan string) {
	ch <- "Message from Go-Routine"
	close(ch)
}
