package main

import (
	"fmt"
	"net"
	"sync"
)

func worker(ports chan int, wg *sync.WaitGroup) {
	for p := range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d", p)
		_, err := net.Dial("tcp", address)
		if err == nil {
			fmt.Println(fmt.Sprintf("%d ***OPENED***", p))
		} else {
			//fmt.Println(fmt.Sprintf("%d closed\n", p))
		}
		wg.Done()
	}
}

func main() {
	ports := make(chan int, 100)
	defer close(ports)

	var wg sync.WaitGroup

	for i := 0; i < cap(ports); i++ {
		go worker(ports, &wg)
	}
	for i := 1; i <= 1024; i++ {
		wg.Add(1)
		ports <- i
	}

	wg.Wait()
}
