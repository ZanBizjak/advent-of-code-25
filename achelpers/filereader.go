package achelpers

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func StringReadFile(file string) string {
	dat, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	return string(dat)
}

func IntReadTwoColumns(filestring string, splitter string) ([]int, []int) {
	file, err := os.Open(filestring)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	col1 := make([]int, 1000)
	col2 := make([]int, 1000)

	for scanner.Scan() {
		row := scanner.Text()
		firstNum, err := strconv.Atoi(strings.Split(row, splitter)[0])
		if err != nil {
			log.Fatal(err)
		}
		secondNum, err := strconv.Atoi(strings.Split(row, splitter)[1])
		if err != nil {
			log.Fatal(err)
		}
		col1 = append(col1, firstNum)
		col2 = append(col2, secondNum)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return col1, col2
}

func IntReadGrid(filestring string, splitter string) [][]int {
	var grid [][]int
	rows := ReadRows(filestring)
	for i, row := range rows {
		grid = append(grid, make([]int, 0))
		strNums := strings.Split(row, splitter)
		for _, strNum := range strNums {
			num, err := strconv.Atoi(strNum)
			if err != nil {
				log.Fatal(err)
			}
			grid[i] = append(grid[i], num)
		}

	}

	return grid
}

func ReadGridRune(filestring string) [][]rune {
	s := ReadRows(filestring)
	var grid [][]rune
	for i, str := range s {
		grid = append(grid, make([]rune, 0))
		grid[i] = []rune(str)
	}

	return grid
}

func ReadRows(filestring string) []string {
	file, err := os.Open(filestring)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var rows []string
	for scanner.Scan() {
		row := scanner.Text()
		rows = append(rows, row)
	}

	return rows
}
