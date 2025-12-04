package day4

import (
	"fmt"

	"github.com/ZanBizjak/advent-of-code-25/achelpers"
)

type Day4 struct {
}

func (Day4) TaskOne() int {
	input := achelpers.ReadGridRune("day4.txt")
	res := 0

	for y := range input {
		for x := range input[y] {
			if input[y][x] == '@' {
				adjPapers := getAdjecentPapers(input, x, y)
				if adjPapers < 4 {
					res += 1
					fmt.Print(string('x'))
				} else {

					fmt.Print(string('@'))
				}
			} else {
				fmt.Print(string('.'))
			}
		}

		fmt.Print("\n")
	}
	return res
}

func getAdjecentPapers(room [][]rune, x, y int) int {
	yLen := len(room)
	xLen := len(room[y])
	numPapers := 0

	if y > 0 {
		numPapers += checkRow(room[y-1], x)
	}
	if y < yLen-1 {
		numPapers += checkRow(room[y+1], x)
	}
	if x > 0 {
		numPapers += checkPoint(room[y][x-1])
	}
	if x < xLen-1 {
		numPapers += checkPoint(room[y][x+1])
	}

	return numPapers
}

func checkPoint(point rune) int {
	if point == '@' {
		return 1
	}
	return 0
}

func checkRow(hall []rune, x int) int {
	numPapers := 0
	hallLen := len(hall)
	if x > 0 && hall[x-1] == '@' {
		numPapers += 1
	}
	if x < hallLen-1 && hall[x+1] == '@' {
		numPapers += 1
	}
	if hall[x] == '@' {
		numPapers += 1
	}

	return numPapers
}

func (Day4) TaskTwo() int {
	room := achelpers.ReadGridRune("day4.txt")
	res := 0
	var resState int
	resState, room = findRemovableRolls(room)
	res += resState

	for resState > 0 {
		resState, room = findRemovableRolls(room)
		res += resState

	}

	return res
}

func findRemovableRolls(room [][]rune) (int, [][]rune) {
	res := 0
	var newRoom [][]rune

	for y := range room {
		newRoom = append(newRoom, make([]rune, 0))
		for x := range room[y] {
			if room[y][x] == '@' {
				adjPapers := getAdjecentPapers(room, x, y)
				if adjPapers < 4 {
					res += 1
					newRoom[y] = append(newRoom[y], '.')
				} else {

					newRoom[y] = append(newRoom[y], '@')
				}
			} else {
				newRoom[y] = append(newRoom[y], '.')
			}
		}
	}

	return res, newRoom
}
