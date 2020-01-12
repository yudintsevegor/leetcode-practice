package longestPalindrome

import (
	"sort"
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

	results := make(Palindromes, 0)
	results = append(results, Palindrome{
		Length: 1,
		Value:  letters[0],
	})

	for {
		if start > length {
			break
		}

		if start == end {
			end = length
			start++
			continue
		}

		// log.Println(letters[start : end+1])
		res, isOk := isPalindrome(letters[start : end+1])
		end--
		if !isOk {
			continue
		}

		// log.Println(res.Value)
		results = append(results, res)
	}
	if len(results) == 0 {
		return ""
	}

	sort.Sort(results)
	/* log.Println("INPUT: ", s)
	for _, r := range results {
		log.Println(r.Value, r.Length)
	}

	*/
	// log.Println(len(results))

	return results[0].Value
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
		// log.Println("HERE", in[i], in[length-i])
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

type Palindromes []Palindrome

func (p Palindromes) Len() int {
	return len(p)
}

func (p Palindromes) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Palindromes) Less(i, j int) bool {
	return p[i].Length > p[j].Length
}
