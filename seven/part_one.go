package seven

import (
	"regexp"

	"github.com/joshsticks/AdventOfCode2020/inputhelper"
)

func recurBags(bags []string, bag string) []string {
	var filteredBags []string
	for _, x := range bags {
		re1 := regexp.MustCompile(`[0-9]\s([a-z]+\s[a-z]+)\sbag`)
		matches := re1.FindAllStringSubmatch(x, -1)

		re2 := regexp.MustCompile(`^([a-z]+\s[a-z]+)\sbags`)
		parent := re2.FindStringSubmatch(x)

		for _, y := range matches {
			if y[1] == bag {
				// fmt.Println(parent[1])
				filteredBags = append(filteredBags, parent[1])
				x := recurBags(bags, parent[1])
				if x != nil {
					filteredBags = append(filteredBags, x...)
				}
			}
		}
	}

	return filteredBags
}

func SolvePart1() int {
	bags := inputhelper.GetStrings("inputs/input7.txt")
	x := recurBags(bags, "shiny gold")
	counter := make(map[string]int)
	for _, row := range x {
		counter[row]++
	}
	return len(counter)
}
