package main

import "fmt"

func main() {
	x := 1
	y := 2
	test(&x, &y)
}

func sumInt(numbers ...int) (sum int, length int) {
	sum = 0
	for _, n := range numbers {
		sum += n
	}

	return len(numbers), sum
}

func minimumFromFour() int {
	var a []int
	var c, m int

	for i := 0; i < 4; i++ {
		fmt.Scan(&c)
		a = append(a, c)
	}

	for i, e := range a {
		if i == 0 || e < m {
			m = e
		}
	}

	return m
}

func vote(x int, y int, z int) int {
	c := x + y + z
	if c >= 2 {
		return 1
	} else {
		return 0
	}
}

/**
Поменять местами переменные на которые ссылаются указатели.
*/
func test(x1 *int, x2 *int) {
	var x int
	x = *x2
	*x2 = *x1
	*x1 = x

	fmt.Print(*x1, *x2)
}
