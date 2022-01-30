package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.TrimFunc(text, func(r rune) bool {
		return !unicode.IsGraphic(r)
	})

	stringLength := utf8.RuneCountInString(text)
	textRunes := []rune(text)
	if stringLength > 0 && unicode.IsUpper(textRunes[0]) && textRunes[stringLength-1] == '.' {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}
}
