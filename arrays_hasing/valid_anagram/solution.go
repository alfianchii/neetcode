package main

import (
	"sort"
	"strings"
)

func sortString(str string) string {
	s := strings.Split(str, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func IsAnagram(s string, t string) bool {
	sortS := sortString(s)
	sortT := sortString(t)

	if sortS != sortT {
		return false
	}

	return true
}
