package main

func IsAnagram(s string, t string) bool {
	mapS := make(map[byte]int)
	mapT := make(map[byte]int)

	for idx, _ := range s {
		if val, exists := mapS[s[idx]]; exists {
			mapS[s[idx]] = val + 1
		} else {
			mapS[s[idx]] = 1
		}
	}

	for idx, _ := range t {
		if val, exists := mapT[t[idx]]; exists {
			mapT[t[idx]] = val + 1
		} else {
			mapT[t[idx]] = 1
		}
	}

	isMapSValid := true
	for idx, _ := range mapS {
		if mapS[idx] != mapT[idx] {
			isMapSValid = false
		}
	}

	isMapTValid := true
	for idx, _ := range mapT {
		if mapS[idx] != mapT[idx] {
			isMapTValid = false
		}
	}

	if isMapSValid && isMapTValid {
		return true
	}

	return false
}
