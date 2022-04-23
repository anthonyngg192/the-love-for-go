package main

import (
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"
)

func main() {
	waitGroup()
	findAddress("web.tryme.com")
	hello()

	http.HandleFunc("/hello", helloworld)
	http.HandleFunc("/headers", headers)
	http.ListenAndServe(":8090", nil)
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
			fmt.Printf("Worker %d starting\n", i)
			time.Sleep(time.Second)
			fmt.Printf("Worker %d done\n", i)
		}()
	}
	wg.Add(5)

	wg.Wait()
}

func findAddress(url string) {
	ips, _ := net.LookupIP(url)
	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func helloworld(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
	fmt.Println(req.Host)
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}
