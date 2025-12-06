package day6

import (
	"strconv"
	"strings"

	"github.com/ZanBizjak/advent-of-code-25/achelpers"
)

type Day6 struct {
}

func (Day6) TaskOne() int {
	input := achelpers.ReadRows("day6.txt")
	cleanedInput := make([][]string, 0)
	for _, row := range input {
		cleanedRow := make([]string, 0)

		for val := range strings.SplitSeq(row, " ") {
			if val != "" {
				cleanedRow = append(cleanedRow, strings.Trim(val, " \n"))
			}
		}
		cleanedInput = append(cleanedInput, cleanedRow)
	}
	numbers := cleanedInput[:len(cleanedInput)-1]
	operators := cleanedInput[len(cleanedInput)-1]

	total := 0
	for i, operator := range operators {
		res := 0
		if operator == "*" {
			res = 1
		}
		for _, numRow := range numbers {
			num, _ := strconv.Atoi(numRow[i])
			if operator == "*" {
				res *= num
				continue
			}
			res += num

		}
		total += res

	}

	return total
}

func (Day6) TaskTwo() int {
	input := achelpers.ReadRows("day6.txt")
	numbers := input[:len(input)-1]
	operators := input[len(input)-1]
	total := 0

	prevFound := 0
	currOperator := operators[0]
	for i := 1; i < len(operators); i++ {
		currRune := operators[i]
		if currRune == ' ' && i != len(operators)-1 {
			continue
		}
		res := 0
		if currOperator == '*' {
			res = 1
		}
		toDeduct := 2
		if i == len(operators)-1 {
			toDeduct = 0
		}
		for j := i - toDeduct; j >= prevFound; j-- {
			numStr := getNumFromCol(numbers, j)
			if numStr == "" {
				continue
			}
			num, _ := strconv.Atoi(numStr)
			if currOperator == '*' {

				res *= num
				continue
			}

			res += num

		}

		currOperator = currRune
		prevFound = i
		total += res

	}

	return total

}

func getNumFromCol(numbers []string, colAt int) string {
	ret := ""
	for _, row := range numbers {
		if row[colAt] == ' ' {
			continue
		}
		ret += string(row[colAt])

	}
	return ret
}
