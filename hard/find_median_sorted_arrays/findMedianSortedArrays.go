package find_median_sorted_arrays

// with nice time but bad memory allocations
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	sumLengths := len(nums1) + len(nums2)
	out := make([]int, 0, sumLengths)

	var (
		i = 0
		j = 0
	)

	for {
		if i > len(nums1)-1 {
			out = append(out, nums2[j:]...)
			break
		}

		if j > len(nums2)-1 {
			out = append(out, nums1[i:]...)
			break
		}

		if nums1[i] <= nums2[j] {
			out = append(out, nums1[i])
			i++
			continue
		}

		if nums1[i] > nums2[j] {
			out = append(out, nums2[j])
			j++
			continue
		}

	}
	if sumLengths%2 != 0 {
		return float64(out[len(out)/2])
	}

	return float64(out[len(out)/2]+out[len(out)/2-1]) / 2
}
