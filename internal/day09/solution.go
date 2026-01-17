package day09

import (
	"fmt"
	"math"
)

// the single line case is edge case and will almost never come true because if 3 points available,
// straight line rectangle will never have the most area.
func Part01(tiles [][2]int) {
	fmt.Println("Solving Part 01")
	ans := 0
	for i := range tiles {
		for j := i + 1; j < len(tiles); j++ {
			// fmt.Printf("Checking tiles:\n\t%v: %v\n\t%v: %v\n", i, tiles[i], j, tiles[j])
			ans = int(math.Max(float64(ans), float64(calcArea(tiles[i], tiles[j]))))
		}
	}

	fmt.Println("Answer: ", ans)
}
func Part02() {}

func calcArea(tile1, tile2 [2]int) int {
	x := int(math.Abs(float64(tile1[0]-tile2[0]))) + 1
	y := int(math.Abs(float64(tile1[1]-tile2[1]))) + 1

	return x * y
}
