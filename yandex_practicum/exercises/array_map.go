package main

import "fmt"

func main() {
	input := []string{
		"cat",
		"dog",
		"bird",
		"dog",
		"parrot",
		"cat",
	}

	fmt.Println(RemoveDuplicates(input))
}

func RemoveDuplicates(input []string) []string {
	res := make([]string, 0, cap(input))
	m := make(map[string]int)
	for _, v := range input {
		if _, ok := m[v]; !ok {
			res = append(res, v)
		}
		m[v]++
	}

	//fmt.Println(m)

	return res
}

func TwoSums() {
	arr := [...]int{2, 7, 11, 15}
	m := make(map[int]int)
	k := 9

	for i, v := range arr {
		if j, ok := m[k-v]; ok {
			fmt.Println(i, j)
		}
		m[v] = i
	}
}
