package main

import (
	"fmt"

	"github.com/joshsticks/AdventOfCode2020/three"
)

func main() {
	// resultOne := one.Solve()
	// fmt.Printf(`problem one: %v`, resultOne)

	// resultTwo := two.SolvePart1()
	// fmt.Printf(`problem two, part1: %v`, resultTwo)

	// resultTwoPartTwo := two.SolvePart2()
	// fmt.Printf(`problem two, part2: %v`, resultTwoPartTwo)

	resultThreePartOne := three.SolvePart1()
	fmt.Printf(`problem three, part1: %v`, resultThreePartOne)

	resultThreePartTwo := three.SolvePart2()
	fmt.Printf(`problem three, part2: %v`, resultThreePartTwo)
}
