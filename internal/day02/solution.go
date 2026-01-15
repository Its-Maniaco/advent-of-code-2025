package day02

import (
	"fmt"
	"strconv"

	"github.com/Its-Maniaco/advent-of-code-2025/internal/utils"
)

func Part01(input []utils.NumRange) {
	fmt.Println("Solving Part 01")

	evenRange := getEvenRange(input)
	ans := 0

	for _, interval := range evenRange {
		n1, n2, diff := 0, 0, 0
		for i := interval.Lower; i <= interval.Upper; i++ {
			nStr := strconv.Itoa(i)
			if nStr[0:len(nStr)/2] == nStr[len(nStr)/2:] {
				if n1 == 0 {
					n1 = i
					ans += n1
					// fmt.Println("\tinvalid: ", n1)
					continue
				}
				if n2 == 0 {
					n2 = i
					// fmt.Println("\tinvalid: ", n2)
					ans += n2
				}
				break
			}
		}
		if n2 != 0 {
			diff = n2 - n1
			// fmt.Printf("  diff between: %v-%v=%v\n", n2, n1, diff)

			insert := n2 + diff
			for insert <= interval.Upper {
				// fmt.Println("\tinvalid insert: ", insert)
				ans += insert
				insert += diff
			}
		}
	}

	fmt.Println("Answer: ", ans)

}

func Part02(input []utils.NumRange) {}
