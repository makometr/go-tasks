package main

import (
	"strconv"
	"strings"
	"unicode/utf8"
)

// объединие двух слайсов

// RunLengthEncode encodes
func RunLengthEncode(str string) string {
	if len(str) == 0 {
		return ""
	}
	var builder strings.Builder
	writeRune := func(char rune, number int) {
		if number != 1 {
			builder.WriteString(strconv.Itoa(number) + string(char))
		} else {
			builder.WriteString(string(char))
		}
	}

	// currentChar := rune(str[0])
	currentChar, _ := utf8.DecodeRuneInString(str)
	counter := 0
	for _, char := range str {
		if char == rune(currentChar) {
			counter++
		} else {
			writeRune(currentChar, counter)
			currentChar = char
			counter = 1
		}
	}
	writeRune(currentChar, counter)

	return builder.String()
}

func main() {
	RunLengthEncode("한한한")
}
