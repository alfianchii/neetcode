package main

import "sort"

func TopKFrequent(nums []int, k int) []int {
	result := make(map[int]int)

	for _, num := range nums {
		result[num]++
	}

	freqs := [][]int{}
	for val, freq := range result {
		freqs = append(freqs, []int{val, freq})
	}
	sort.Slice(freqs, func(i, j int) bool {
		return freqs[i][1] > freqs[j][1]
	})

	res := make([]int, 0, k)
	for i := 0; i < k; i++ {
		res = append(res, freqs[i][0])
	}

	return res
}
