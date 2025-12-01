package day1

import (
	"log"
	"math"
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

			res += int(math.Abs(math.Floor(float64(dialAt) / 100.0)))
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
			res += int(math.Abs(math.Floor(float64(dialAt) / 100.0)))
		}

		dialAt = mod(dialAt, 100)

	}
	return res
}

func mod(a, b int) int {
	return (a%b + b) % b
}
