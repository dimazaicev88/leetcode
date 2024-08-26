package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(maximumNumberOfStringPairs([]string{"cd", "ac", "dc", "ca", "zz"}))
}

//func maximumNumberOfStringPairs(words []string) int {
//	var count int
//	mapWords := make(map[string]bool)
//
//	for i := 0; i < len(words); i++ {
//		if _, ok := mapWords[words[i]]; ok {
//			count++
//			continue
//		}
//		mapWords[Reverse(words[i])] = true
//	}
//	return count
//}

func maximumNumberOfStringPairs(words []string) int {
	var count int
	arrWords := make([]string, len(words), len(words))

	for i := 0; i < len(words); i++ {
		arrWords[i] = Reverse(words[i])
	}

	for i := 0; i < len(words)-1; i++ {
		if slices.Contains(arrWords[i:len(words)], words[i]) {
			count++
			continue
		}
	}
	return count
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
