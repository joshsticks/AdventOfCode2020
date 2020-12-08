package eight

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/joshsticks/AdventOfCode2020/inputhelper"
)

func runProgram2(program []string) int {
	accumulator := 0
	visited := make(map[int]bool)
	i := 0
	re := regexp.MustCompile(`([a-z]{3})\s(\+|\-)([0-9]+)`)

	for !visited[i] && i < len(program) {
		visited[i] = true
		instruction := re.FindStringSubmatch(program[i])
		switch command := instruction[1]; command {
		case "acc":
			val, _ := strconv.Atoi(instruction[3])
			if instruction[2] == "+" {
				accumulator += val
			} else {
				accumulator -= val
			}
			i++
		case "jmp":
			val, _ := strconv.Atoi(instruction[3])
			if instruction[2] == "+" {
				i += val
			} else {
				i -= val
			}
		case "nop":
			i++
		}
	}

	fmt.Printf("program is this long %v, it stopped at %v\n", len(program), i)
	if i == len(program) {
		return accumulator
	}
	return 0
}

func GenerateAllProgramPermutations(program []string) [][]string {
	var newPrograms [][]string
	for i := 0; i < len(program); i++ {
		if strings.Contains(program[i], "jmp") || strings.Contains(program[i], "nop") {
			newProgram := make([]string, len(program))
			copy(newProgram, program)
			replacer := strings.NewReplacer("jmp", "nop", "nop", "jmp")
			newProgram[i] = replacer.Replace(newProgram[i])
			newPrograms = append(newPrograms, newProgram)
		}
	}
	return newPrograms
}

func SolvePart2() int {
	program := inputhelper.GetStrings("inputs/input8.txt")
	newPrograms := GenerateAllProgramPermutations(program)
	for _, x := range newPrograms {
		val := runProgram2(x)
		if val > 0 {
			return val
		}
	}
	return 0
}
