package five

import (
	"math"

	"github.com/joshsticks/AdventOfCode2020/inputhelper"
)

func FindRow(x string) int {
	rowMin := 0.0
	rowMax := 127.0

	letters := []rune(x)
	for i := 0; i < 7; i++ {
		if string(letters[i]) == "F" {
			if i == 6 {
				return int(rowMin)
			}

			rowMax = math.Floor((rowMax + rowMin) / 2)
		} else if string(letters[i]) == "B" {
			if i == 6 {
				return int(rowMax)
			}

			rowMin = math.Ceil((rowMax + rowMin) / 2)
		}

	}
	return 0
}

func FindColumn(x string) int {
	colMin := 0.0
	colMax := 7.0

	letters := []rune(x)
	for i := 7; i < 10; i++ {
		if string(letters[i]) == "L" {
			if i == 9 {
				return int(colMin)
			}

			colMax = math.Floor((colMax + colMin) / 2)
		} else if string(letters[i]) == "R" {
			if i == 9 {
				return int(colMax)
			}

			colMin = math.Ceil((colMax + colMin) / 2)
		}

	}
	return 0
}

func FindSeatID(x string) int {
	return FindRow(x)*8 + FindColumn(x)
}

func SolvePart1() int {

	seats := inputhelper.GetStrings("inputs/input5.txt")

	maxSeatID := 0

	for _, x := range seats {
		seatID := FindSeatID(x)
		if seatID > maxSeatID {
			maxSeatID = seatID
		}
	}

	return maxSeatID
}
