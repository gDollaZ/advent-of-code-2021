package day1

import (
	"advent-of-code-2021/utils"
	"fmt"
)

func Run(inputFile string) {
	fmt.Println("Day 1!")
	depthArray := utils.FileToIntSlice("day1/day1.txt")
	fmt.Println("  Part 1: ", partOne(depthArray))
	fmt.Println("  Part 2: ", partTwo(depthArray))
}

func partTwo(depthArray []int) int {
	var newArray []int
	for i, v := range depthArray {
		if i == len(depthArray)-2 {
			break
		}
		newVal := v + depthArray[i+1] + depthArray[i+2]
		newArray = append(newArray, newVal)
	}
	return partOne(newArray)
}

func partOne(depthArray []int) int {
	increaseCount := 0
	for i, v := range depthArray {
		if i == 0 {
			continue
		}

		previous := depthArray[i-1]
		if previous < v {
			increaseCount += 1
		}
	}
	return increaseCount
}
