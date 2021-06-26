package task34

import "unicode"

// Написать программу, которая проверяет, что все символы в строке уникальные.

// Task34 is
func Task34(s string) bool {
	unique := make(map[rune]bool, len(s))

	for _, r := range s {
		if unicode.IsSpace(r) {
			continue
		}
		if _, ok := unique[r]; ok {
			return false
		}
		unique[r] = true
	}

	return true
}
