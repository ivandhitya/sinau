package main

import (
	"fmt"
	"net/http"
	"time"
)

func handleRequest(w http.ResponseWriter, r *http.Request, ch chan string) {
	// Proses panjang
	time.Sleep(3 * time.Second)
	ch <- "Request Processed 1"
}

func handler(w http.ResponseWriter, r *http.Request) {
	ch := make(chan string)
	go handleRequest(w, r, ch)

	result := <-ch
	fmt.Fprintf(w, result)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is running on port 9999")
	if err := http.ListenAndServe(":9999", nil); err != nil {
		fmt.Println(err)
	}
}
