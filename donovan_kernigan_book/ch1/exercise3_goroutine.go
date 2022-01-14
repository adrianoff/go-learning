package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	urls := [2]string{"https://facebook.com", "https://google.com"}

	for _, url := range urls {
		go fetch(url, ch)
	}
	for range urls {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	//start := time.Now()
	_, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	ch <- url + " DONE"
	//_, err := httputil.DumpResponse(resp, true)
	//fmt.Println(string(b))

	//nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	//resp.Body.Close() // don't leak resources
	//if err != nil {
	//	ch <- fmt.Sprintf("while reading %s: %v", url, err)
	//	return
	//}
	//secs := time.Since(start).Seconds()
	//ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
