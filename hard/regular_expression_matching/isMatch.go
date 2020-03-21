package regular_expression_matching

import (
	"log"
	"reflect"
	"strings"
)

func isMatch(s string, p string) bool {
	str := strings.Split(s, "")
	pattern := strings.Split(p, "")

	if len(str) == 0 && len(pattern) == 0 {
		return true
	} else if !strings.Contains(p, "*") && !strings.Contains(p, "*") && len(str) != len(pattern) {
		return false
	}

	var (
		sInd   = 0
		pInd   = 0
		result = make([]string, 0, len(str))
	)

	for {
		if sInd > len(str)-1 {
			if pInd < len(pattern)-1 {
				if pattern[pInd] != "*" {
					result = append(result, pattern[pInd:]...)
					break
				}
				result = append(result, str[sInd-1])
			}
			break
		} else if pInd > len(pattern)-1 {
			break
		}

		switch {
		case pattern[pInd] == str[sInd]:
			result = append(result, str[sInd])
			pInd++
			sInd++
			continue
		case pattern[pInd] == ".":
			result = append(result, str[sInd])
			pInd++
			sInd++
			continue
		case pattern[pInd] == "*":
			previous := pattern[pInd-1]
			if true {
				// TODO
			} else if previous == str[sInd] {
				result = append(result, str[sInd])
				sInd++
				continue
			} else if previous == "." {
				result = append(result, str[sInd])
				sInd++
				continue
			}
			pInd++
		case pattern[pInd] != str[sInd]:
			pInd++
			continue
		}
	}

	log.Printf("input string: `%v`; input pattern: `%v`; result: `%v`", s, p, strings.Join(result, ""))

	return reflect.DeepEqual(result, str)
}
