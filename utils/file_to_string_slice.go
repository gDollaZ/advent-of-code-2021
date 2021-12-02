package utils

import (
	"bufio"
	"log"
	"os"
)

func FileToStringSlice(fileName string) []string {
	var stringSlice []string
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i := scanner.Text()
		stringSlice = append(stringSlice, i)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return stringSlice
}
