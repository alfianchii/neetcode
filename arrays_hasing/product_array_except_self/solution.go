package main

func ProductExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	count := 1

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			num := nums[j]

			if i == j {
				continue
			}

			count *= num
		}

		res[i] = count
		count = 1
	}

	return res
}
