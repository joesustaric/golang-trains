//Package io contains functions to parse input for trains app
package io

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

//ParseInput returns an array of station connections in the format of AB5
func ParseInput(filePath string) []string {
	file := openFile(filePath)
	defer file.Close()
	result := readLines(file)
	return createTrimedArray(result)
}

func readLines(file *os.File) []string {
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	check(scanner.Err())
	result := strings.Split(lines[0], ",")
	return result
}

func openFile(filePath string) *os.File {
	file, err := os.Open(filePath)
	check(err)
	return file
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func createTrimedArray(input []string) []string {
	result := make([]string, len(input))
	for index, element := range input {
		result[index] = strings.TrimSpace(element)
	}
	return result
}

func validateInput(connection string) bool {
	result, _ := regexp.MatchString("^[A-Za-z]{2}[0-9]+", connection)
	return result
}
