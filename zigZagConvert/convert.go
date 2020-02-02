package zigZagConvert

import (
	"strings"
)

func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}

	letters := strings.Split(s, "")

	mapHelper := make(map[int][]string, numRows)
	for i := range mapHelper {
		mapHelper[i] = make([]string, 0, 1)
	}

	isDownWalk := true
	nextRow := 0

	for i, letter := range letters {
		ind := getInd(&isDownWalk, &nextRow, i, numRows)
		// log.Print(ind)
		mapHelper[ind] = append(mapHelper[ind], letter)
	}

	out := ""
	for i := 0; i < numRows; i++ {
		out += strings.Join(mapHelper[i], "")
	}
	return out
}

func getInd(isDownWalk *bool, nextRow *int, currentInd, rowsNumber int) int {
	rowsMinusOne := rowsNumber - 1
	if currentInd < rowsMinusOne {
		*nextRow++
		return currentInd
	}

	if currentInd%rowsMinusOne == 0 {
		isNewDownWalk := *isDownWalk
		if isNewDownWalk {
			*nextRow--
			*isDownWalk = !isNewDownWalk
			return rowsMinusOne
		}
		*isDownWalk = !isNewDownWalk
		*nextRow++
		return 0
	}

	if *isDownWalk {
		*nextRow++
		return currentInd % rowsMinusOne
	}

	newInd := *nextRow
	*nextRow--

	return newInd
}
