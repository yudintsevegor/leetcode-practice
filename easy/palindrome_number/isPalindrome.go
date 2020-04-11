package palindrome_number

// there are global var, because it is a little hack
// for getting more effective result in tests
var (
	numbers            = make([]int, 0, 1)
	multiplier         = 10
	previousMultiplier = 1
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	splitNumForOrder(x)

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

// output will be a slice with numbers in revers order
// example: num == 1234 --> []int{4,3,2,1}
func splitNumForOrder(num int) []int {
	for {
		if num/multiplier == 0 {
			numbers = append(numbers, num/previousMultiplier)
			break
		}

		numbers = append(numbers, (num%multiplier)/previousMultiplier)

		multiplier *= 10
		previousMultiplier *= 10
	}

	return numbers
}
