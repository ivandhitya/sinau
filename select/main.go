package main

import (
	"fmt"
	"time"
)

func sendMessage(ch chan string, sleep time.Duration, msg string) {
	time.Sleep(sleep)
	ch <- "Message from Channel " + msg
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go sendMessage(ch1, time.Second*3, "1")
	go sendMessage(ch2, time.Second*2, "2")
loop:
	for {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		case <-time.After(4 * time.Second):
			break loop
		}
	}
}
