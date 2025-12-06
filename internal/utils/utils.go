package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// Open input or test file using flag.
// 0 for input file else 1 for test file
func OpenInput(fileLoc string, isTestFile bool) (*os.File, error) {
	// if fileLoc == "" {
	// 	return nil, fmt.Errorf("missing file location")
	// }

	if isTestFile {
		fileLoc += "input.txt"
	} else {
		fileLoc += "test.txt"
	}

	file, err := os.Open(fileLoc)
	if err != nil {
		return nil, err
	}

	return file, nil
}

// returns a basic terminal input reader
func GetTerminalInputReader() *bufio.Reader {
	return bufio.NewReader(os.Stdin)
}

// Get number of digits in input number
func GetNumberOfDigits(a int) int {
	if a == 0 {
		return 1
	}

	count := 0

	for a != 0 {
		a /= 10
		count++
	}

	return count
}

// This function assumes that the line only contains integer numbers and whitespace
// and skips any whitespace encountered
func StringLineToIntSlice(line string) ([]int, error) {
	a := []int{}
	fields := strings.Fields(line)

	for _, f := range fields {
		t, err := strconv.Atoi(f)
		if err != nil {
			return nil, err
		}

		a = append(a, t)
	}

	return a, nil
}
