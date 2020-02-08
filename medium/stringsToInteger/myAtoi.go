package stringsToInteger

import (
	"math"
	"strconv"
	"strings"
)


func myAtoi(str string) int {
	if len(str) == 0{
		return 0
	}

	splitted := strings.Split(strings.TrimSpace(str), "")
	if len(splitted) == 0{
		return 0
	}

	var allowedSymbols = map[string]struct{}{
		"+": {},
		"-": {},
		"0": {},
		"1": {},
		"2": {},
		"3": {},
		"4": {},
		"5": {},
		"6": {},
		"7": {},
		"8": {},
		"9": {},

	}

	if _, ok := allowedSymbols[splitted[0]]; !ok{
		return 0
	}

	plusSign := "+"
	minusSign := "-"


	if splitted[0] != plusSign && splitted[0] != minusSign{
		delete(allowedSymbols, plusSign)
		delete(allowedSymbols, minusSign)
	}

	out := ""
	counts := 0
	for _, symbol := range splitted{
		if _, ok := allowedSymbols[symbol]; !ok{
			break
		}

		if symbol == plusSign || symbol == minusSign{
			counts++
			if counts > 1{
				break
			}
		}

		out += symbol
	}

	res, err := strconv.ParseFloat(out,64)
	if err != nil{
		return 0
	}

	edge := math.Pow(2,31)
	if math.Abs(res) >=  edge{
		if res < 0{
			return -int(edge)
		}

		return int(edge-1)
	}

	return int(res)
}
