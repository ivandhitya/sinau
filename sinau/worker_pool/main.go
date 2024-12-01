package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- string) {
	for j := range jobs { // mempeorses pekerjaan/ job yang datang
		fmt.Printf("Worker %d mulai pekerjaan %d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("Worker %d selesai pekerjaan %d\n", id, j)
		hasilPekerjaan := j * 2

		// mengirimkan hasil pekerjaan ke channel result
		results <- fmt.Sprintf("Worker %d hasil pekerjaan %d adalah %d \n", id, j, hasilPekerjaan)
	}
}

func main() {
	jobs := make(chan int, 5)
	results := make(chan string, 5)

	// membuat 3 worker
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// mengirimkan 5 pekerjaan kepada fungsi worker melalui channel "job"
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// menampilkan hasil pemrosesan dari fungsi worker melalui channel result
	for a := 1; a <= 5; a++ {
		fmt.Println(<-results)
	}

	time.Sleep(10 * time.Second)
}
