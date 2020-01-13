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

	maxSlice := letters[:1]
	maxLen := len(maxSlice)

	for {
		if start > length || maxLen > (length-start) {
			break
		}

		if start == end || maxLen > (end-start) {
			end = length
			start++
			continue
		}

		slice := letters[start : end+1]
		end--

		isOk := isPalindrome(slice)
		if !isOk {
			continue
		}

		sliceLen := len(slice)

		if sliceLen > maxLen {
			maxSlice = slice
			maxLen = sliceLen
		}
	}

	return strings.Join(maxSlice, "")
}

func isPalindrome(in []string) bool {
	length := len(in) - 1
	if length == 0 {
		return true
	}

	for i := 0; i < length/2+1; i++ {
		if in[i] != in[length-i] {
			return false
		}
	}

	return true
}
