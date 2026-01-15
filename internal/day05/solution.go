package day05

import (
	"fmt"

	"github.com/Its-Maniaco/advent-of-code-2025/internal/utils"
)

func Part01(rangeStr []string, nums []int) {
	fmt.Println("Solving Part 01")
	nR := utils.ConvertStringRange(rangeStr)
	merged := utils.UnifyRange(nR)
	ans := 0
	for _, num := range nums {
		left, right := 0, len(merged)-1

		for left <= right {
			mid := left + (right-left)/2
			if merged[mid].Lower <= num && merged[mid].Upper >= num {
				ans++
				break
			} else if merged[mid].Lower > num {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}

	fmt.Println("Answer: ", ans)
}

func Part02(rangeStr []string, nums []int) {
	fmt.Println("Solving Part 02")
	nR := utils.ConvertStringRange(rangeStr)
	merged := utils.UnifyRange(nR)
	ans := 0
	for _, interval := range merged {
		ans += interval.Upper - interval.Lower + 1
	}

	fmt.Println("Answer: ", ans)
}
