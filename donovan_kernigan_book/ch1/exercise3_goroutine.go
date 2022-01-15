package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	urls := []string{"https://facebook.com", "https://google.com", "https://ya.ru", "https://youtube.com"}

	for _, url := range urls {
		go fetch(url, ch)
	}

	var a string
	for {
		_, _ = fmt.Fscan(os.Stdin, &a)
		if a == "STOP" {
			break
		}
		fmt.Println(a)
		ch <- a
	}

	for {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	ch <- url + " DONE"
	_, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	ch <- url + " DONE"

}

func readFromUser(ch chan<- string) {
	var a string
	for {
		_, _ = fmt.Fscan(os.Stdin, &a)
		ch <- a
	}
}
