package longestPalindrome

import (
	"strings"
)

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return s
	}

	letters := strings.Split(s, "")
	if len(letters) == 1 {
		return letters[0]
	}

	var (
		length = len(letters) - 1
		start  = 0
		end    = length
	)

	maxElement := Palindrome{
		Length: 1,
		Value:  letters[0],
	}

	for {
		if start > length {
			break
		}

		if start == end || maxElement.Length > (end-start) {
			end = length
			start++
			continue
		}

		res, isOk := isPalindrome(letters[start : end+1])
		end--
		if !isOk {
			continue
		}

		if res.Length > maxElement.Length {
			maxElement = res
		}
	}

	return maxElement.Value
}

type Palindrome struct {
	Length int
	Value  string
}

func isPalindrome(in []string) (Palindrome, bool) {
	length := len(in) - 1
	if length == 0 {
		return Palindrome{
			Length: 1,
			Value:  in[0],
		}, true
	}

	for i := 0; i < length/2+1; i++ {
		if in[i] != in[length-i] {
			return Palindrome{}, false
		}
	}

	p := Palindrome{
		Value:  strings.Join(in, ""),
		Length: length + 1,
	}

	return p, true
}
