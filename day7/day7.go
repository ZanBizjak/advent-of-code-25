package day7

import (
	"log"

	"github.com/ZanBizjak/advent-of-code-25/achelpers"
)

type Day7 struct {
}

func (Day7) TaskOne() int {
	input := achelpers.ReadGridRune("day7.txt")
	start := locateStart(input[0])
	res := 0
	var beamLocs map[int]bool
	prevBeamLocs := make(map[int]bool)
	prevBeamLocs[start] = true

	grid := input[2:]
	if start == -1 {
		log.Fatal("Can't find start :(")
	}
	for _, row := range grid {
		beamLocs = make(map[int]bool)
		for beam, _ := range prevBeamLocs {
			if row[beam] == '^' {
				res += 1
				prevBeamLocs[beam] = false
				if beam > 0 {
					beamLocs[beam-1] = true
				}
				if beam < len(row)-1 {
					beamLocs[beam+1] = true
				}
			}
		}
		for beam, didntSplit := range prevBeamLocs {
			if didntSplit {
				beamLocs[beam] = true
			}
		}

		prevBeamLocs = beamLocs

	}

	return res
}

func locateStart(row []rune) int {
	for i, r := range row {
		if r == 'S' {
			return i
		}
	}
	return -1

}

func (Day7) TaskTwo() int {
	input := achelpers.ReadGridRune("day7.txt")
	startX := locateStart(input[0])
	startY := 0
	savedResults := make(map[int]map[int]int)
	res := findAllTimelines(input, startX, startY, savedResults)
	return res
}

func findAllTimelines(grid [][]rune, currX, currY int, savedResults map[int]map[int]int) int {
	if savedResults[currX][currY] != 0 {
		return savedResults[currX][currY]
	}
	if len(savedResults[currX]) == 0 {
		savedResults[currX] = make(map[int]int)
	}
	if currY+1 == len(grid) {
		savedResults[currX][currY] = 1
		return 1
	}
	nextRow := grid[currY+1]
	if nextRow[currX] == '.' {
		savedResults[currX][currY] = findAllTimelines(grid, currX, currY+1, savedResults)
	} else if currX < 1 {
		savedResults[currX][currY] = 1 + findAllTimelines(grid, currX+1, currY+1, savedResults)
	} else if currX == len(nextRow)-1 {
		savedResults[currX][currY] = findAllTimelines(grid, currX-1, currY+1, savedResults) + 1
	} else {

		savedResults[currX][currY] = findAllTimelines(grid, currX-1, currY+1, savedResults) + findAllTimelines(grid, currX+1, currY+1, savedResults)
	}

	return savedResults[currX][currY]

}
