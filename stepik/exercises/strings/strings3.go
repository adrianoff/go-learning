package main

import (
	"fmt"
	"strings"
)

func main() {
	var str, substr string

	_, _ = fmt.Scan(&str)
	_, _ = fmt.Scan(&substr)

	fmt.Print(strings.Index(str, substr))
}
