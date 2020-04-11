package integer_to_roman

import (
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	splited := strings.Split(s, "")
	length := len(splited)
	count := 0

	for i := 0; i < length; i++ {
		chars := make(map[string]struct{})
		for j := i; j < length; j++ {
			if _, ok := chars[splited[j]]; ok {
				break
			}

			chars[splited[j]] = struct{}{}
			count++
		}

		if count > maxLen {
			maxLen = count
		}

		count = 0
	}

	return maxLen
}
