package day2

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/ZanBizjak/advent-of-code-25/achelpers"
)

type Day2 struct {
}

func (Day2) TaskOne() int {
	input := achelpers.StringReadFile("day2.txt")
	res := 0
	for idRanges := range strings.SplitSeq(input, ",") {
		firstId, lastId := getSplitIds(idRanges)
		res += findAllInvalidIdsHalved(firstId, lastId)
	}

	return res
}

func (Day2) TaskTwo() int {

	input := achelpers.StringReadFile("day2.txt")
	res := 0
	for idRanges := range strings.SplitSeq(input, ",") {
		firstId, lastId := getSplitIds(idRanges)
		sumOfInvalidIds, err := findAllInvalidIds(firstId, lastId)
		if err != nil {
			log.Panic(fmt.Errorf("Error when trying to get sum of invalid ids in range %s: %w", idRanges, err))
		}
		res += sumOfInvalidIds
	}

	return res
}

func findAllInvalidIds(firstId, lastId int) (int, error) {
	res := 0
	for i := firstId; i <= lastId; i++ {
		if i < 10 {
			continue
		}
		id := strconv.FormatInt(int64(i), 10)
		isRepeating, err := hasRepeatingPattern(id)
		if err != nil {
			return -1, fmt.Errorf("Error when finding out if %s has a repeating pattern: %w", id, err)
		}
		if isRepeating {
			res += int(i)
		}

	}

	return res, nil
}

func hasRepeatingPattern(id string) (bool, error) {
	length := len(id)
	for i := 1; i <= length/2; i++ {
		if length%i != 0 {
			continue
		}
		subIds, err := makeSubIdsByLength(id, i)
		if err != nil {
			return false, fmt.Errorf("Error when creating subIds for id %s with denominator of %d: %w", id, i, err)
		}
		if isRepeating(subIds) {

			return true, nil
		}

	}

	return false, nil
}

func isRepeating(subIds []int) bool {
	var prevSubId int
	for i, subId := range subIds {
		if i == 0 {
			prevSubId = subId
			continue
		}
		if prevSubId != subId {
			return false
		}
		prevSubId = subId
	}

	return true
}

func makeSubIdsByLength(id string, subIdSize int) ([]int, error) {
	subIds := make([]int, 0, len(id)/subIdSize)
	for i := 0; i < len(id); i += subIdSize {
		subId, err := strconv.Atoi(id[i : i+subIdSize])
		if err != nil {
			return nil, fmt.Errorf("Failed to convert string %s to integer of length %d: %w", id[i:], subIdSize, err)
		}
		subIds = append(subIds, subId)
	}
	return subIds, nil
}

func findAllInvalidIdsHalved(firstId, lastId int) int {
	res := 0
	for i := firstId; i <= lastId; i++ {
		id := strconv.FormatInt(int64(i), 10)
		idlen := len(id)
		if idlen%2 != 0 {
			continue
		}
		firstHalf, secondHalf := id[0:idlen/2], id[idlen/2:]
		if firstHalf == secondHalf {
			res += int(i)
		}

	}

	return res
}

func getSplitIds(idRanges string) (int, int) {
	splitIds := achelpers.StrToIntSlice(idRanges, "-")
	return splitIds[0], splitIds[1]

}
