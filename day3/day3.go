package day3

import (
	"advent-of-code-2021/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Run(iFile string) {
	fmt.Println("Day 3!")
	diagnosticReport := utils.FileToStringSlice(iFile)
	mostCommonValues := getCommonValues("most", diagnosticReport)
	leastCommonValues := getCommonValues("least", diagnosticReport)
	gamma := convertBinary(mostCommonValues)
	epsilon := convertBinary(leastCommonValues)
	fmt.Println("  Part 1: ", gamma*epsilon)
}

func convertBinary(binary string) int64 {
	val, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return val
}

func getWantedValues(values map[int]map[string]int, mode string) []string {
	wantedValues := make([]string, 12)
	for i := 0; i < len(values); i++ {
		wantedBit := "0"
		for bit, count := range values[i] {
			if bit == "0" {
				continue
			}
			switch mode {
			case "most":
				if count > values[i][wantedBit] {
					wantedBit = bit
				}
			case "least":
				if count < values[i][wantedBit] {
					wantedBit = bit
				}
			}
		}
		wantedValues[i] = wantedBit
	}
	return wantedValues
}

func getCommonValues(mode string, diagnosticReport []string) string {
	var commonValues string
	valuesCounts := make(map[int]map[string]int)
	for index, bits := range diagnosticReport {
		bitsSplit := strings.Split(bits, "")
		for i, bit := range bitsSplit {
			if index == 0 {
				valuesCounts[i] = map[string]int{"1": 0}
				valuesCounts[i]["0"] = 0
			}
			valuesCounts[i][bit] += 1
		}
	}
	commonValues = strings.Join(getWantedValues(valuesCounts, mode), "")
	return commonValues
}
