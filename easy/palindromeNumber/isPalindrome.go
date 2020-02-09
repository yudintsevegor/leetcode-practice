package palindromeNumber

var (
	numbers            = make([]int, 0, 1)
	multiplier         = 10
	previousMultiplier = 1
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	for {
		if x/multiplier == 0 {
			numbers = append(numbers, x/previousMultiplier)
			break
		}

		numbers = append(numbers, (x%multiplier)/previousMultiplier)

		multiplier *= 10
		previousMultiplier *= 10
	}

	multiplier = 10
	previousMultiplier = 1

	length := len(numbers)
	if length < 2 {
		numbers = numbers[:0]
		return true
	}

	for i := 0; i < length/2; i++ {
		if numbers[i] != numbers[length-1-i] {
			numbers = numbers[:0]
			return false
		}
	}

	numbers = numbers[:0]
	return true
}
