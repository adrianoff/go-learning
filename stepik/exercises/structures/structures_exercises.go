package main

import "fmt"

type Circle struct {
	x int
	y int
}

func (c *Circle) sum() int {
	return c.x + c.y
}

func main() {
	c := structureTest()
	fmt.Print(c)

	c2 := Circle{1, 2}
	fmt.Print(c2)

	fmt.Println(c2.sum())
}

func structureTest() Circle {
	c := Circle{}
	return c
}
