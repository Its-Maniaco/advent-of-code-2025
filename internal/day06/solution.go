package day06

import (
	"fmt"
	"log"

	"github.com/Its-Maniaco/advent-of-code-2025/internal/utils"
)

func Part01(lines []string) {
	fmt.Println("Solving Part 01")
	ans := int64(0)
	// 0th to calculate plus while 1st index to calcukate multiplication
	calcArr := [][2]int{}
	for j, line := range lines {
		// fmt.Println("Reading line: ", line)
		// run till 2nd last line so we can parse numbers
		if j < len(lines)-1 {
			lineInt, err := utils.StringLineToIntSlice(line)
			if err != nil {
				log.Fatal("error getting int from string")
			}

			for i, num := range lineInt {
				// fmt.Println("Checking num:", num)
				// fmt.Printf("\n\tCalcArr: %v, on line: %v\n", calcArr, lineInt)
				if i >= len(calcArr) {
					calcArr = append(calcArr, [2]int{num, num})
					continue
				}

				calcArr[i][0] = calcArr[i][0] + num
				calcArr[i][1] = calcArr[i][1] * num

				// fmt.Println("CalcArr now: ", calcArr)
			}
		} else {
			colNum := 0
			for _, sign := range line {
				switch sign {
				case ' ':
					continue
				case '+':
					ans += int64(calcArr[colNum][0])
					colNum++
				case '*':
					ans += int64(calcArr[colNum][1])
					colNum++
				default:
					continue
				}
			}
		}

	}

	fmt.Println("Answer is: ", ans)
}
