package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	for i := 1; i <= 1024; i++ {
		wg.Add(1)
		go func(j int) {
			address := fmt.Sprintf("scanme.nmap.org:%d", j)
			fmt.Printf("Scaning %d.", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				fmt.Printf(" %d closed.\n", j)
				return
			}

			conn.Close()
			fmt.Printf(" %d ***OPENED***\n", j)
		}(i)
	}
}
