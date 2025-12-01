package day1

import (
	"log"
	"strconv"
	"strings"

	"github.com/ZanBizjak/advent-of-code-25/achelpers"
)

type Day1 struct {
}

func (d Day1) TaskOne() int {
	input := achelpers.ReadRows("day1.txt")
	dialAt := 50
	res := 0

	for _, i := range input {
		side := i[0]
		numStr := strings.ReplaceAll(i, "L", "")
		numStr = strings.ReplaceAll(numStr, "R", "")
		num, err := strconv.Atoi(numStr)
		if err != nil {
			log.Fatal(err)
		}
		if side == 'L' {
			dialAt -= num
		}
		if side == 'R' {
			dialAt += num
		}
		if dialAt < 0 || dialAt > 99 {
			dialAt = mod(dialAt, 100)
		}
		if dialAt == 0 {

			res += 1
		}

	}
	return res
}

func (d Day1) TaskTwo() int {
	input := achelpers.ReadRows("day1.txt")
	dialAt := 50
	res := 0

	for _, i := range input {
		dialStart := dialAt
		side := i[0]
		numStr := strings.ReplaceAll(i, "L", "")
		numStr = strings.ReplaceAll(numStr, "R", "")
		num, err := strconv.Atoi(numStr)
		if err != nil {
			log.Fatal(err)
		}
		wholeRotations := num / 100
		num -= wholeRotations * 100
		if side == 'L' {
			res += handleLeft(dialStart, dialAt, num)
			dialAt -= num
		}
		if side == 'R' {
			res += handleRight(dialAt, num)
			dialAt += num
		}
		res += wholeRotations
		dialAt = mod(dialAt, 100)
	}
	return res
}

func handleRight(dialAt, num int) int {
	dial := dialAt + num
	if dial > 99 {
		return 1
	}
	return 0

}

func handleLeft(dialStart, dialAt, num int) int {
	dial := dialAt - num
	if dial <= 0 && dialStart != 0 {
		return 1
	}
	return 0
}

func mod(a, b int) int {
	return (a%b + b) % b
}
