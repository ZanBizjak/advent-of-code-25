package day3

import (
	"log"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/ZanBizjak/advent-of-code-25/achelpers"
)

type Day3 struct {
}

func (Day3) TaskOne() int {
	input := achelpers.ReadRows("day3.txt")
	res := 0
	for _, row := range input {
		row = strings.Trim(row, "\n ")
		joltage := findJoltage(row)
		res += joltage
		log.Printf("In %s found joltage of %d\n", row, joltage)
	}
	return res
}

func findJoltage(row string) int {
	bank := achelpers.StrToIntSlice(row, "")
	firstBiggest, firstIndex := findBiggestFrom(bank, 0, true)
	secondBiggest, secondIndex := findBiggestFrom(bank, firstIndex, false)
	if firstIndex < secondIndex {
		firstBiggest *= 10
	} else {
		secondBiggest *= 10
	}

	return int(firstBiggest + secondBiggest)
}

func findBiggestFrom(bank []int, i int, fromStart bool) (int, int) {
	ignoreLast := 0
	if i == len(bank)-1 || fromStart {
		i = 0
		if !fromStart {
			ignoreLast = 1
		}
	} else {
		i += 1
	}
	currBiggest := bank[i]
	currIndex := i
	for ; i < len(bank)-ignoreLast; i++ {
		if bank[i] > currBiggest {
			currBiggest = bank[i]
			currIndex = i
		}
	}

	return currBiggest, currIndex

}

func (Day3) TaskTwo() int {
	input := achelpers.ReadRows("day3.txt")
	res := 0
	for _, row := range input {
		row = strings.Trim(row, "\n ")
		joltage := findJoltageTwelve(row)
		res += joltage
		// log.Printf("In %s found joltage of %d\n", row, joltage)
	}
	return res
}

func findJoltageTwelve(row string) int {
	bank := achelpers.StrToIntSlice(row, "")
	foundJolts := make([]int, 0, 12)
	leftMostJolt := -1
	for len(foundJolts) < 12 {
		biggestIndex := findBiggestFromIgnoringIndicies(bank, foundJolts, leftMostJolt)
		if leftMostJolt == -1 || leftMostJolt > biggestIndex {
			leftMostJolt = biggestIndex
		}

		foundJolts = append(foundJolts, biggestIndex)
	}

	sort.Slice(foundJolts, func(i, j int) bool {
		return foundJolts[j] > foundJolts[i]

	})
	joltage := ""
	for i := range len(foundJolts) {
		joltage += strconv.Itoa(bank[foundJolts[i]])
	}
	intJoltage, err := strconv.Atoi(joltage)
	if err != nil {
		panic(err)
	}

	return intJoltage
}

func findBiggestFromIgnoringIndicies(bank []int, ignoreIndicies []int, from int) int {

	currIndex := -1
	startPoint := 0
	if len(bank)-from-len(ignoreIndicies)-1 > 0 {
		startPoint = from + 1
	}
	// 0 1 2 3 4 5 6 7 8 9 10 11 12 13 14
	//
	// 2 3 4 2 3 4 2 3 4 2  3  4  2  7  8
	// 	 4   3 4   3 4    3  4     7  8

	// len(bank) {15} - from + 1 {3} = 13
	// len(bank) {15} - from + 1 {3} - len(ignoreIndices) {9} = 4
	for i := startPoint; i < len(bank); i++ {
		if slices.Contains(ignoreIndicies, i) {
			continue
		}
		if currIndex == -1 {
			currIndex = i
			continue
		}
		if bank[i] >= bank[currIndex] {
			currIndex = i
		}
	}

	return currIndex

}
