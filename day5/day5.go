package day5

import (
	"fmt"
	"log"
	"strconv"

	"github.com/ZanBizjak/advent-of-code-25/achelpers"
)

type Day5 struct {
}

func (Day5) TaskOne() int {
	input := achelpers.ReadRows("day5.txt")
	res := 0
	elIndex := findEL(input)
	ranges := input[:elIndex]
	ids := input[elIndex+1:]
	for _, id := range ids {
		res += isBetween(id, ranges)
	}

	return res
}

func isBetween(id string, ranges []string) int {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	for _, fromTo := range ranges {
		from, to := getSplitIds(fromTo)
		if idInt >= from && idInt <= to {
			log.Print(idInt, " is fresh, because range ", fromTo)
			return 1
		}
	}
	return 0
}

func getSplitIds(idRanges string) (int, int) {
	splitIds := achelpers.StrToIntSlice(idRanges, "-")
	return splitIds[0], splitIds[1]

}

func findEL(input []string) int {
	for i, line := range input {
		if line == "" {
			return i
		}
	}
	panic("COULDN'T FIND EL")
}

func (Day5) TaskTwo() int {
	input := achelpers.ReadRows("day5.txt")
	elIndex := findEL(input)
	ranges := input[:elIndex]
	res := 0
	for i, fromTo := range ranges {
		adjustedRange := adjustRange(fromTo, i, ranges)
		ranges[i] = adjustedRange

	}

	for _, fromTo := range ranges {
		if fromTo == "invalid" {
			continue
		}
		from, to := getSplitIds(fromTo)
		res += to - from + 1
	}

	return res
}

func adjustRange(fromTo string, start int, ranges []string) string {
	from, to := getSplitIds(fromTo)
	for i, span := range ranges {
		if span == "invalid" {
			continue
		}
		if i == start {
			continue
		}
		fromCmp, toCmp := getSplitIds(ranges[i])
		if (from >= fromCmp && from <= toCmp) && (to >= fromCmp && to <= toCmp) {
			return "invalid"
		}
		if to >= fromCmp && to <= toCmp {
			to = fromCmp - 1
		}
		if from >= fromCmp && from <= toCmp {
			from = toCmp + 1
		}
		if to < from {
			return "invalid"
		}
	}

	ret := fmt.Sprintf("%d-%d", from, to)
	return ret

}

func appendAll(setOfItems map[int]bool, from, to int) map[int]bool {
	for i := from; i <= to; i++ {
		setOfItems[i] = true
	}
	return setOfItems
}
