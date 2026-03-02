package main

func LongestCommonPrefix(strs []string) string {
	currStr := ""
	for i, str := range strs {
		if i == 0 {
			currStr = str
			continue
		}

		for j, ch := range currStr {
			if j == len(str) || j == len(currStr) {
				currStr = str[:j]
				break
			}

			if byte(ch) != str[j] {
				currStr = str[:j]
				break
			}
		}
	}

	return currStr
}
