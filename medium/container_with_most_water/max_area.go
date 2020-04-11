package container_with_most_water

func maxArea(heights []int) int {
	maxVolume := 0

	for ind, height := range heights {
		if ind == len(heights)-1 {
			continue
		}

		difference := len(heights) - 1 - ind
		for internalInd := len(heights) - 1; internalInd > ind; internalInd-- {
			currentVolume := getMin(height, heights[internalInd]) * difference
			if maxVolume < currentVolume {
				maxVolume = currentVolume
			}
			difference--
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
