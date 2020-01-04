package twoSum

func twoSum(nums []int, target int) []int {
	for i := range nums {
		for k := i + 1; k < len(nums); k++ {
			if nums[i]+nums[k] == target {
				return []int{i, k}
			}
		}
	}

	return []int{}
}
