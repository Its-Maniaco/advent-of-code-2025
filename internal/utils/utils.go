package utils

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

type NumRange struct {
	Lower int
	Upper int
}

// convert string range to integer range
func ConvertStringRange(rangeStr []string) []NumRange {
	nR := []NumRange{}

	for _, v := range rangeStr {
		vInts := strings.Split(v, "-")
		lowerStr := vInts[0]
		lower, err := strconv.Atoi(lowerStr)
		if err != nil {
			fmt.Printf("error converting lower: %v\n", err)
			return nil
		}
		upperStr := vInts[1]
		upper, err := strconv.Atoi(upperStr)
		if err != nil {
			fmt.Printf("error converting upper: %v\n", err)
			return nil
		}

		nR = append(nR, NumRange{Lower: lower, Upper: upper})
	}

	return nR
}

func UnifyRange(nR []NumRange) []NumRange {
	if len(nR) == 0 {
		return nil
	}
	// first we sort the range based on lower limit
	slices.SortFunc(nR, func(a, b NumRange) int {
		if a.Lower != b.Lower {
			return a.Lower - b.Lower
		}
		// tiebreaker
		return a.Upper - b.Upper
	})

	// now we merge
	merged := []NumRange{}
	curr := nR[0]

	for i := 1; i < len(nR); i++ {
		next := nR[i]
		if next.Lower <= curr.Upper+1 {
			if next.Upper > curr.Upper {
				curr.Upper = next.Upper
			}
		} else {
			merged = append(merged, curr)
			curr = next
		}
	}

	merged = append(merged, curr)
	return merged
}
