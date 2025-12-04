package main

import (
	"errors"
	"log"
	"os"

	"github.com/ZanBizjak/advent-of-code-25/day1"
	"github.com/ZanBizjak/advent-of-code-25/day2"
	"github.com/ZanBizjak/advent-of-code-25/day3"
	"github.com/ZanBizjak/advent-of-code-25/day4"
)

func main() {
	day := os.Args[2]
	task := os.Args[3]

	solver, err := whichAdventDay(day)
	if err != nil {
		log.Fatal(err)
	}

	if task == "1" {
		log.Print(solver.TaskOne())
		return
	}

	if task == "2" {
		log.Print(solver.TaskTwo())
		return
	}

}

func whichAdventDay(day string) (AdventDay, error) {

	switch day {
	case "1":
		return day1.Day1{}, nil
	case "2":
		return day2.Day2{}, nil
	case "3":
		return day3.Day3{}, nil
	case "4":
		return day4.Day4{}, nil
	default:
		return nil, errors.New("Can't find solver for this day: " + day)
	}
}
