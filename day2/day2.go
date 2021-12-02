package day2

import (
	"advent-of-code-2021/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Run(iFile string) {
	commands := utils.FileToStringSlice(iFile)
	fmt.Println("Day 2!")
	pilot(commands)
}

func pilot(commands []string) {
	position := []int{0, 0}
	position2 := []int{0, 0}
	aim := 0
	for _, command := range commands {
		direction := strings.Split(command, " ")[0]
		stepCount, err := strconv.Atoi(strings.Split(command, " ")[1])
		if err != nil {
			log.Fatal(err)
		}

		switch direction {
		case "forward":
			position[0] += stepCount
			position2[0] += stepCount
			position2[1] += aim * stepCount
		case "down":
			position[1] += stepCount
			aim += stepCount
		case "up":
			position[1] -= stepCount
			aim -= stepCount
		default:
			fmt.Println("Direction not recognized")
		}
	}
	fmt.Println("  Part 1: ", position[1]*position[0])
	fmt.Println("  Part 2: ", position2[1]*position2[0])
}
