package day05

import "fmt"

func Part01(rangeStr []string, nums []int) {
	fmt.Println("Solving Part 01")
	nR := convertStringRange(rangeStr)
	merged := unifyRange(nR)
	ans := 0
	for _, num := range nums {
		left, right := 0, len(merged)-1

		for left <= right {
			mid := left + (right-left)/2
			if merged[mid].lower <= num && merged[mid].upper >= num {
				ans++
				break
			} else if merged[mid].lower > num {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}

	fmt.Println("Answer: ", ans)
}

func Part02() {}
