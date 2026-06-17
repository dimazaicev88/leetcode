package _205

import "fmt"

func isIsomorphic(s string, t string) bool {
	mapS := make(map[rune]rune)
	mapT := make(map[rune]rune)
	for i := 0; i < len(s); i++ {
		_, existsS := mapS[rune(t[i])]
		_, existsT := mapT[rune(s[i])]

		if existsS || existsT {
			fmt.Println(mapS[rune(t[i])], mapT[rune(s[i])])
			if mapS[rune(t[i])] != mapT[rune(s[i])] {
				return false
			}
		} else {
			fmt.Println(rune(s[i]), rune(t[i]))
			mapS[rune(s[i])] = rune(t[i])
			mapT[rune(t[i])] = rune(s[i])
		}
	}

	return true
}
