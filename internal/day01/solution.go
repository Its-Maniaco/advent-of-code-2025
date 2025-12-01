package day01

import "fmt"

func Part01(input []int) {
	last := 50
	current := last
	zeroCount := 0
	for _, num := range input {
		current += num
		if current != 0 {
			if current < 0 {
				for current < 0 {
					current += 100
				}
			} else if current > 0 {
				for current >= 100 {
					current -= 100
				}
			}
		}

		if current == 0 {
			zeroCount++
		}

	}

	fmt.Println("Number of zero's: ", zeroCount)
}
