package day3

import (
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
	}
	return res
}

func findMax(bank []int) (int, int) {
	maxIndex := -1

	for i := range bank {
		if maxIndex == -1 {
			maxIndex = i
			continue
		}

		if bank[i] > bank[maxIndex] {
			maxIndex = i
		}
	}

	return maxIndex, bank[maxIndex]

}

func findJoltageTwelve(row string) int {
	bank := achelpers.StrToIntSlice(row, "")

	foundJolts := make([]int, 0)
	for len(foundJolts) < 12 {
		leftBank := bank[:len(bank)-11+len(foundJolts)]
		leftMostIndex, leftMostValue := findMax(leftBank)
		foundJolts = append(foundJolts, leftMostValue)
		bank = bank[leftMostIndex+1:]
		if len(bank)+len(foundJolts) == 12 {
			for len(bank) != 0 {
				foundJolts = append(foundJolts, bank[0])
				bank = bank[1:]
			}
		}

	}

	joltage := ""
	for i := range len(foundJolts) {
		joltage += strconv.Itoa(foundJolts[i])
	}
	intJoltage, err := strconv.Atoi(joltage)
	if err != nil {
		panic(err)
	}

	return intJoltage
}
