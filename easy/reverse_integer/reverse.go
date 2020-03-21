package reverse_integer

import (
	"math"
	"strconv"
	"strings"
)

func reverse(x int) int {
	edgePower := math.Pow(2, 31)
	if edgePower < math.Abs(float64(x)) {
		return 0

	}

	splited := strings.Split(strconv.Itoa(x), "")
	if x < 0 {
		splited = splited[1:]
	}

	length := len(splited) - 1
	for i := 0; i <= (length)/2; i++ {
		fromLeft := splited[i]
		splited[i] = splited[length-i]
		splited[length-i] = fromLeft
	}

	joined := strings.Join(splited, "")
	res, _ := strconv.Atoi(joined)

	if edgePower < math.Abs(float64(res)) {
		return 0

	}

	if x < 0 {
		res = -res
	}

	return res

}
