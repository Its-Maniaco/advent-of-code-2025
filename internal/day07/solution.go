package day07

import (
	"fmt"

	"github.com/Its-Maniaco/advent-of-code-2025/internal/utils"
)

func Part01(lines []string) {
	ans := 0
	// currline

	// prevMap will store position of beam in the previous line
	// currMap will store the position of beams found in line being read
	prevBeam, currBeam := utils.NewSet[int](), utils.NewSet[int]()

	for i, line := range lines {
		// fmt.Println("Reading line: ", line)
		if i == 0 {
			for j, beam := range line {
				if beam == 'S' {
					// assuming a single starting beam
					prevBeam.Add(j)
					break
				}
			}
			continue
		}
		for b := range prevBeam { // b beam location
			if line[b] == '^' {
				ans++
				if b-1 >= 0 && b-1 <= len(line)-1 {
					currBeam.Add(b - 1)
				}
				if b+1 >= 0 && b+1 <= len(line)-1 {
					currBeam.Add(b + 1)
				}
			} else {
				currBeam.Add(b)
			}
		}

		// prevBeam = currBeam
		// currBeam = nil // panics: map can be nil; nil map can be read from but cannot be written to

		prevBeam, currBeam = currBeam, utils.NewSet[int]()

	}

	fmt.Println("Answer: ", ans)
}
