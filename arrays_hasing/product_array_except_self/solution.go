package main

func ProductExceptSelf(nums []int) []int {
	res := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		count := 1
		for j := 0; j < len(nums); j++ {
			if i != j {
				count *= nums[j]
			}
		}

		res[i] = count
	}

	return res
}
