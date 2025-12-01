package utils

import (
	"bufio"
	"os"
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
