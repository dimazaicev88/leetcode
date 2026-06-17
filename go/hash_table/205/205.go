package _205

func isIsomorphic(s string, t string) bool {
	mapS := make(map[rune]rune)
	mapT := make(map[rune]rune)

	for i := 0; i < len(s); i++ {
		runeS := rune(s[i])
		runeT := rune(t[i])
		valS, existsS := mapS[runeS]
		valT, existsT := mapT[runeT]

		if existsS {
			if valS != runeT {
				return false
			}
		} else if existsT {
			if valT != runeS {
				return false
			}
		} else {
			mapS[runeS] = runeT
			mapT[runeT] = runeS
		}
	}

	return true
}
