package five

import (
	"fmt"
	"sort"

	"github.com/joshsticks/AdventOfCode2020/inputhelper"
)

func GetMissingNos(a []int) []int {
	var missingNos []int
	for i := 0; i < len(a); i++ {
		if i > 0 {
			if a[i] != a[i-1]+1 {
				missingNos = append(missingNos, a[i]-1)
			}
		}
	}
	return missingNos
}

func SolvePart2() {

	seats := inputhelper.GetStrings("inputs/input5.txt")

	maxSeatID := 0

	var seatIDs []int

	for _, x := range seats {
		seatID := FindSeatID(x)
		seatIDs = append(seatIDs, seatID)
		if seatID > maxSeatID {
			maxSeatID = seatID
		}
	}

	sort.Ints(seatIDs)
	// for _, x := range seatIDs {
	// 	fmt.Println(x)
	// }
	missingNos := GetMissingNos(seatIDs)
	for _, x := range missingNos {
		fmt.Println(x)
	}
}
