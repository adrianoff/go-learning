package main

import "fmt"

func main() {
	for i := 1; i < 100; i++ {

		switch {
		case i % 3 == 0 && i % 5 == 0 :
			fmt.Print("FizzBuzz\n")
		case i % 3 == 0:
			fmt.Print("Fizz\n")
		case i % 5 == 0:
			fmt.Print("Buzz\n")
		default:
			fmt.Printf("%d\n", i)
		}
	}
}
