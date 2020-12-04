package four

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/joshsticks/AdventOfCode2020/inputhelper"
)

func getValue(x string, y string) {

}

func validateYearField(x string, min int, max int, field string) bool {
	reg := fmt.Sprintf(`%v:([0-9]*)`, field)
	re := regexp.MustCompile(reg)
	match := re.FindStringSubmatch(x)
	if match != nil {
		year, _ := strconv.Atoi(match[1])

		if year >= min && year <= max {
			return true
		}
	}

	return false
}

func validateHgt(x string) bool {
	re := regexp.MustCompile("hgt:([0-9]*)(cm|in)")
	match := re.FindStringSubmatch(x)
	if match != nil {
		hgt, _ := strconv.Atoi(match[1])
		if match[2] == "cm" {
			if hgt >= 150 && hgt <= 193 {
				return true
			}
		} else if match[2] == "in" {
			if hgt >= 59 && hgt <= 76 {
				return true
			}
		}
	}
	return false
}

func validateHcl(x string) bool {
	re := regexp.MustCompile(`hcl:#[0-9a-f]{6}`)
	match := re.FindStringSubmatch(x)
	if match != nil {
		return true
	}
	return false
}

func validateEcl(x string) bool {
	re := regexp.MustCompile(`ecl:(amb|blu|brn|gry|grn|hzl|oth)`)
	match := re.FindStringSubmatch(x)
	if match != nil {
		return true
	}
	return false
}

func validatePid(x string) bool {
	re := regexp.MustCompile(`pid:([0-9]{9})([^0-9]|$)`)
	match := re.FindStringSubmatch(x)
	if match != nil {
		return true
	}
	return false
}

func SolvePart2() int {

	passports := inputhelper.GetStringsFromFileByDelim("inputs/input4.txt")

	count := 0

	for _, x := range passports {
		if validateYearField(x, 1920, 2002, "byr") && validateYearField(x, 2010, 2020, "iyr") && validateYearField(x, 2020, 2030, "eyr") && validateHgt(x) && validateHcl(x) && validateEcl(x) && validatePid(x) {
			count++
		}
	}

	return count
}
