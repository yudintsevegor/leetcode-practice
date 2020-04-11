package integer_to_roman

import (
	"strconv"
	"strings"
)

const (
	I = "I" // 1
	V = "V" // 5
	X = "X" // 10
	L = "L" // 50
	C = "C" // 100
	D = "D" // 500
	M = "M" // 1000
)

func intToRoman(num int) string {
	numbers := splitNumForOrderWithLoop(num)
	var result string

	for i := len(numbers) - 1; i >= 0; i-- {
		switch i {
		case 0:
			result += linkToRoman(numbers[i], I, I+V, V, I+X)
		case 1:
			result += linkToRoman(numbers[i], X, X+L, L, X+C)
		case 2:
			result += linkToRoman(numbers[i], C, C+D, D, C+M)
		case 3:
			result += strings.Repeat("M", numbers[i])
		}
		// in task we have range between 1 and 3999
		// as a result we dont have more than 4 orders
	}

	return result
}

func linkToRoman(num int, one, four, five, nine string) string {
	if num < 4 {
		return strings.Repeat(one, num)
	} else if num == 4 {
		return four
	} else if num == 5 {
		return five
	} else if 5 < num && num < 9 {
		return five + strings.Repeat(one, num-5)
	} else if num == 9 {
		return nine
	}

	return ""
}

// output will be a slice with numbers in revers order
// example: num == 1234 --> []int{4,3,2,1}
func splitNumForOrderWithLoop(num int) []int {
	var (
		numbers            = make([]int, 0, 1)
		multiplier         = 10
		previousMultiplier = 1
	)

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

// more readable version, but less effective
func splitNumForOrderWithoutLoop(num int) []int {
	numbers := strings.Split(strconv.Itoa(num), "")
	out := make([]int, 0, len(numbers))

	for i := len(numbers) - 1; i >= 0; i-- {
		convertedNum, _ := strconv.Atoi(numbers[i])
		out = append(out, convertedNum)
	}
	return out
}
