package _25

import (
	"unicode"
)

func isPalindrome(s string) bool {
	if s == " " {
		return true
	}

	var cleanStr []rune

	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			cleanStr = append(cleanStr, unicode.ToLower(r))
		}
	}

	left, right := 0, len(cleanStr)-1

	for left < right {
		if cleanStr[left] != cleanStr[right] {
			return false
		}

		left++
		right--
	}

	return true
}
