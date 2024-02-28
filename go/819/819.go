package _819

import (
	"regexp"
	"slices"
	"strings"
)

func mostCommonWord(paragraph string, banned []string) string {
	word := ""
	cnt := 0
	m := regexp.MustCompile("[.,?]")
	paragraph = m.ReplaceAllString(paragraph, "")
	splitStr := strings.Split(paragraph, " ")
	mapString := make(map[string]int)
	for _, str := range splitStr {
		if slices.Contains(banned, str) == false {
			mapString[strings.ToLower(str)]++
		}
	}

	for key, value := range mapString {
		if value > cnt {
			word = key
			cnt = value
		}

	}
	if cnt == 1 {
		return strings.ToLower(splitStr[0])
	}
	return strings.ToLower(word)
}
