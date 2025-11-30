package achelpers

import (
	"log"
	"strconv"
	"strings"
)

func IntRemoveIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func RuneCopySlice(s []rune) []rune {
	ret := make([]rune, 0)
	return append(ret, s...)
}

func IntCopySlice(s []int64) []int64 {
	ret := make([]int64, 0)
	return append(ret, s...)
}

func RuneCopyGrid(s [][]rune) [][]rune {
	ret := make([][]rune, 0)
	for _, arr := range s {
		ret = append(ret, RuneCopySlice(arr))
	}

	return ret
}

func StrToIntSlice(s string, seperator string) []int64 {
	var ret []int64
	for _, sNum := range strings.Split(s, seperator) {
		num, err := strconv.ParseInt(sNum, 10, 64)
		if err != nil {
			log.Fatal("could not convert element of slice to int")
		}
		ret = append(ret, num)
	}

	return ret

}
