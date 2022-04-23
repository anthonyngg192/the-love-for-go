package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	waitGroup()
}

func hello() {
	var wg sync.WaitGroup
	fmt.Println(wg)

	wg.Add(1)

	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		j, more := <-jobs
		if more {
			fmt.Println("received job", j)
		} else {
			fmt.Println("received all job")
			done <- true
			return
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}

	close(jobs)
	fmt.Println("send all jobs")
}

var wg sync.WaitGroup

func waitGroup() {

	for i := 1; i <= 5; i++ {
		i := i
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}
	wg.Add(1)

	wg.Wait()
}

func worker(i int) {
	fmt.Printf("Worker %d starting\n", i)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", i)
}
