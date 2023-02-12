package main

import (
	"fmt"
	"net"
)

func main() {
	for i := 1; i <= 1024; i++ {
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
