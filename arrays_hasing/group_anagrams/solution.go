package main

import (
	"slices"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	freq := make(map[rune]int)

	for _, ch := range s {
		freq[ch]++
	}

	for _, ch := range t {
		freq[ch]--
		if freq[ch] < 0 {
			return false
		}
	}

	return true
}

func GroupAnagrams(strs []string) [][]string {
	result := [][]string{}
	for i := 0; i < len(strs); i++ { // 1
		// ALL STRS [act pots tops cat stop hat]
		// STRS [act, pots, stop, hat]
		s := strs[i] // pots
		// RESULT [[act cat], [pots, tops]]
		result = slices.Insert(result, len(result), []string{s})
		for j := i + 1; j < len(strs); j++ { // 1 + 1 = 2
			t := strs[j] // tops
			if ok := isAnagram(s, t); ok {
				result[i] = slices.Insert(result[i], len(result[i]), t)
				strs = slices.Delete(strs, j, j+1)
				j--
			}
		}
	}

	return result
}
