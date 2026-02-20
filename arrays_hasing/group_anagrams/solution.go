package main

import (
	"sort"
)

func GroupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)

	for _, s := range strs {
		chars := []byte(s)

		sort.Slice(chars, func(i, j int) bool {
			return chars[i] < chars[j]
		})

		key := string(chars)
		groups[key] = append(groups[key], s)
	}

	res := make([][]string, 0, len(groups))
	for _, group := range groups {
		res = append(res, group)
	}

	return res
}
