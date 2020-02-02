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
		ind, newCurrentRow, newDownWalk := getInd(isDownWalk, i, nextRow, numRows)
		// log.Print(ind)
		mapHelper[ind] = append(mapHelper[ind], letter)

		nextRow = newCurrentRow
		isDownWalk = newDownWalk
	}

	out := ""
	for i := 0; i < numRows; i++ {
		out += strings.Join(mapHelper[i], "")
	}
	return out
}

func getInd(isDownWalk bool, currentInd, nextRow, rowsNumber int) (int, int, bool) {
	rowsMinusOne := rowsNumber - 1
	if currentInd < rowsMinusOne {
		nextRow++
		return currentInd, nextRow, isDownWalk
	}

	if currentInd%rowsMinusOne == 0 {
		if isDownWalk {
			nextRow--
			return rowsNumber - 1, nextRow, !isDownWalk
		}
		nextRow++
		return 0, nextRow, !isDownWalk
	}

	if isDownWalk {
		nextRow++
		return currentInd % rowsMinusOne, nextRow, isDownWalk
	}

	newInd := nextRow
	nextRow--

	return newInd, nextRow, isDownWalk
}
