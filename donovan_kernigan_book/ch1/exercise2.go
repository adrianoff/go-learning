package main

import (
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	for _, arg := range os.Args[1:] {
		counts[arg]++
	}

	fmt.Println(counts)
}
