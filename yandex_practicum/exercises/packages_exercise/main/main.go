package main

import (
	"fmt"
	"exercise"
)

func main() {
	//mathslices.SumSlice(make([]int, 0))
	fmt.Println(mathslices.SumSlice([]int{1, 4, 6}))

	slice := []int{1, 4, 6}
	mathslices.MapSlice(slice, func(i int) int {
		return i*2
	})
	fmt.Println(slice)
}


