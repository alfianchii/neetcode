package main

func ProductExceptSelf(nums []int) []int {
	res := make([]int, len(nums))

	for i := range nums {
		count := 1
		for j := range nums {
			if i != j {
				count *= nums[j]
			}
		}

		res[i] = count
	}

	return res
}
