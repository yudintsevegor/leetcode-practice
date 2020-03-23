package container_with_most_water

func maxArea(heights []int) int {
	maxVolume := 0

	for ind, height := range heights {
		if ind == len(heights)-1 {
			continue
		}

		difference := 1
		for internalInd := ind + 1; internalInd < len(heights); internalInd++ {
			volume := difference * height
			if volume < maxVolume {
				difference++
				continue
			}

			currentVolume := getMin(volume, heights[internalInd]*difference)
			if maxVolume < currentVolume {
				maxVolume = currentVolume
			}
			difference++
		}
	}

	return maxVolume
}

func getMin(first, second int) int {
	if first < second {
		return first
	}
	return second
}
