package day05

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type numRange struct {
	lower int
	upper int
}

// convert string range to integer range
func convertStringRange(rangeStr []string) []numRange {
	nR := []numRange{}

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

		nR = append(nR, numRange{lower: lower, upper: upper})
	}

	return nR
}

func unifyRange(nR []numRange) []numRange {
	if len(nR) == 0 {
		return nil
	}
	// first we sort the range based on lower limit
	slices.SortFunc(nR, func(a, b numRange) int {
		if a.lower != b.lower {
			return a.lower - b.lower
		}
		// tiebreaker
		return a.upper - b.upper
	})

	// now we merge
	merged := []numRange{}
	curr := nR[0]

	for i := 1; i < len(nR); i++ {
		next := nR[i]
		if next.lower <= curr.upper+1 {
			if next.upper > curr.upper {
				curr.upper = next.upper
			}
		} else {
			merged = append(merged, curr)
			curr = next
		}
	}

	merged = append(merged, curr)
	return merged
}
