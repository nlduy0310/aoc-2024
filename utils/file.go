package utils

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLines(file string) ([]string, error) {
	fileRef, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("an error occurred when opening file \"%s\": %w", file, err)
	}
	defer fileRef.Close()

	ret := make([]string, 0)
	scanner := bufio.NewScanner(fileRef)
	for scanner.Scan() {
		ret = append(ret, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("an error occurred when reading file \"%s\": %w", file, err)
	}

	return ret, nil
}

func MustReadLines(fileName string) []string {
	lines, err := ReadLines(fileName)
	PanicIf(err)

	return lines
}
