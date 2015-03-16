//Package io contains functions to parse input for trains app
package io

import (
	"bufio"
	"os"
	"strings"
)

//ParseInput returns an array of station connections in the format of AB5
func ParseInput(filePath string) ([]string, error) {

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	result := strings.Split(lines[0], ",")

	x := make([]string, len(result))

	for index, element := range result {
		x[index] = strings.TrimSpace(element)
	}
	return x, scanner.Err()

}
