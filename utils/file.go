package utils

import (
	"bufio"
	"os"
)

func MustReadLines(fileName string) []string {

	file, err := os.Open(fileName)
	PanicIf(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	ret := make([]string, 0)

	for scanner.Scan() {
		ret = append(ret, scanner.Text())
	}

	PanicIf(scanner.Err())

	return ret
}
