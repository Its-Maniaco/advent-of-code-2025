package day04

import "fmt"

type point struct {
	i int
	j int
}

func Part01(input []string, paper byte) {
	fmt.Println("Solving Part 01")
	ans := 0
	for i, line := range input {
		for j, ele := range line {
			if ele != rune(paper) {
				continue
			}
			count := 0
			up, down, left, right := i-1, i+1, j-1, j+1
			// upper line is present
			if up >= 0 {
				// check diagonal top-left
				if left >= 0 && input[up][left] == paper {
					count++
				}
				// check above
				if input[up][j] == paper {
					count++
				}
				// check diagonal top-right
				if right < len(line) && input[up][right] == paper {
					count++
				}
			}
			// front and back element
			if left >= 0 && input[i][left] == paper {
				count++
			}
			if right < len(line) && input[i][right] == paper {
				count++
			}
			// check if we already exceeded the count
			if count >= 4 {
				continue
			}
			// bottom line is present
			if down < len(input) {
				// check diagonal bottom-left
				if left >= 0 && input[down][left] == paper {
					count++
				}
				// check bottom
				if input[down][j] == paper {
					count++
				}
				// check diagonal bottom-right
				if right < len(line) && input[down][right] == paper {
					count++
				}
			}

			if count < 4 {
				ans++
			}
		}
	}

	fmt.Println("Answer: ", ans)
}

func Part02(input []string, paper byte) {
	mp := make(map[point]bool)
	fmt.Println("Solving Part 02")
	ans := 0
	for i, line := range input {
		for j, ele := range line {
			if ele != rune(paper) {
				continue
			}
			mp[point{i, j}] = true
			count := 0
			up, down, left, right := i-1, i+1, j-1, j+1
			if up >= 0 {
				if left >= 0 && mp[point{up, left}] == true {
					count++
				}
				if mp[point{up, j}] == true {
					count++
				}
				if right < len(line) && mp[point{up, right}] == true {
					count++
				}
			}
			if left >= 0 && mp[point{i, left}] == true {
				count++
			}
			if right < len(line) && line[right] == paper {
				count++
			}

			if down < len(input) {
				// check diagonal bottom-left
				if left >= 0 && input[down][left] == paper {
					count++
				}
				// check bottom
				if input[down][j] == paper {
					count++
				}
				// check diagonal bottom-right
				if right < len(line) && input[down][right] == paper {
					count++
				}
			}
			if count < 4 {
				ans++
				delete(mp, point{i, j})
			}
		}
	}
	// printFinal(input, mp)

	// now we iterate over map to recheck remaining ones
	flag := false
	for !flag {
		mpLenStart := len(mp)
		for k, v := range mp {
			// fmt.Printf("\t%v,%v\n", k.i, k.j)
			if v != true {
				delete(mp, k)
				continue
			}
			count := 0
			up, down, left, right := k.i-1, k.i+1, k.j-1, k.j+1
			// check above row
			if up >= 0 {
				if left >= 0 && mp[point{up, left}] == true {
					count++
				}
				if mp[point{up, k.j}] == true {
					count++
				}
				if right < len(input[0]) && mp[point{up, right}] == true {
					count++
				}
			}

			// check row
			if left >= 0 && mp[point{k.i, left}] == true {
				count++
			}
			if right < len(input[0]) && mp[point{k.i, right}] == true {
				count++
			}

			// check bottom row
			if down < len(input) {
				if left >= 0 && mp[point{down, left}] == true {
					count++
				}
				if mp[point{down, k.j}] == true {
					count++
				}
				if right < len(input[0]) && mp[point{down, right}] == true {
					count++
				}
			}
			// fmt.Println("count: ", count)
			if count < 4 {
				ans++
				delete(mp, point{k.i, k.j})
			}

		}
		if len(mp) == mpLenStart {
			flag = true
		}
	}
	printFinal(input, mp)

	fmt.Println("Answer: ", ans)
}

func printFinal(input []string, mp map[point]bool) {
	fmt.Println("--------------------------------------------")
	for i, line := range input {
		for j, ele := range line {
			if ele == '@' && mp[point{i, j}] {
				fmt.Print(string(ele))
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

	fmt.Println("--------------------------------------------")
}
