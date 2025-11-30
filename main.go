package main

import (
	"errors"
	"log"
	"os"

	"zan.bizjak/aoc-25/day1"
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
		log.Print("YAY")
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
	default:
		return nil, errors.New("Can't find solver for this day: " + day)
	}
}
