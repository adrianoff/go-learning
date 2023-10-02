package mathslices

func SumSlice(s []int) int {

	sum := 0
	for _, v := range s {
		sum = sum+v
	}

	return sum
}

func MapSlice(s []int, op func(int) int) {
	for i, v := range s {
		s[i] = op(v)
	}
}

func FoldSlice(s []int, op func(int, int) int, init int) int {
	return 0
}
