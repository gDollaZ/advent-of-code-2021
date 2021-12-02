package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func FileToIntSlice(fileName string) []int {
	var intSlice []int
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		intSlice = append(intSlice, i)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return intSlice
}
