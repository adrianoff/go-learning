package main

import "fmt"

func main() {
	m := map[string]int{
		"bread":  50,
		"milk":   100,
		"butter": 200,
	}

	for k, v := range m {
		if v > 70 {
			fmt.Println(k)
		}
	}

	fmt.Println(m)
}
