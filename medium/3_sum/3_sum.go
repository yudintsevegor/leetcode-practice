package __sum

import (
	"sort"
	"strconv"
)

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	} else if len(nums) == 3 {
		sum := 0
		for _, v := range nums {
			sum += v
		}
		if sum != 0 {
			return nil
		}
		return [][]int{nums}
	}

	mapNums := make(map[int][]int, len(nums)) // key-->value: num-->[]indexes
	for i, num := range nums {
		if v, ok := mapNums[num]; ok {
			mapNums[num] = append(v, i)
			continue
		}
		mapNums[num] = []int{i}
	}

	out := make([][]int, 0, 1024)
	uniq := make(map[string]int)
	for i := 0; i <= len(nums)-3; i++ {
		for j := i + 1; j <= len(nums)-2; j++ {
			indexes, ok := mapNums[0-(nums[i]+nums[j])]
			if !ok {
				continue
			}

			k := -1
			for searchedIND, ind := range indexes {
				if ind <= j || ind <= i {
					continue
				}
				k = ind
				indexes = append(indexes[:searchedIND], indexes[searchedIND+1:]...)
				break
			}

			if k < 0 {
				continue
			}

			arr := []int{nums[i], nums[j], nums[k]}
			sort.Ints(arr)

			key := strconv.Itoa(arr[0]) + strconv.Itoa(arr[1]) + strconv.Itoa(arr[2])
			if uniq[key] == 0 {
				uniq[key]++
				out = append(out, arr)
			}
		}
	}

	return out
}
