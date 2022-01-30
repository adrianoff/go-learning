package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.TrimFunc(text, func(r rune) bool {
		return !unicode.IsGraphic(r)
	})
	text = strings.ReplaceAll(text, " ", "")

	runes := []rune(text)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			fmt.Print("Нет")
			os.Exit(0)
		}
	}

	fmt.Print("Палиндром")
}
