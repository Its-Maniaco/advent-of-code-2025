package day03

import (
	"fmt"
	"log"
	"math"

	"github.com/Its-Maniaco/advent-of-code-2025/internal/utils"
)

func Part01(input []string) {
	fmt.Println("Solving Part 01")
	if len(input) == 0 {
		return
	}

	ans := 0

	for _, s := range input {
		if len(s) < 2 {
			continue
		} else if len(s) == 2 {
			ans += int(s[0]-'0')*10 + int(s[1]-'0')
		}

		// Find first largest digit (leftmost)
		firstVal := int(s[0] - '0')
		firstPos := 0
		for i := 1; i < len(s)-1; i++ {
			n := int(s[i] - '0')
			if n > firstVal {
				firstVal = n
				firstPos = i
			}
		}

		// Find largest digit after firstVal
		secondVal := int(s[firstPos+1] - '0')
		for i := firstPos + 1; i < len(s); i++ {
			n := int(s[i] - '0')
			if n > secondVal {
				secondVal = n
			}
		}

		ans += firstVal*10 + secondVal
	}

	fmt.Println("Ans: ", ans)
}

func Part02(input []string) {
	fmt.Println("Solving Part 02")
	if len(input) == 0 {
		return
	}

	ans := int64(0)

	for _, s := range input {
		// fmt.Println("Starting process for: ", s)
		stack := utils.NewStack()
		for k := 0; k < len(s); k++ {
			valK := int(s[k] - '0')
			// fmt.Println("	Current Stack ", stack)
			// fmt.Println("		Checking: ", valK)
			if stack.IsEmpty() {
				stack = stack.Push(valK)
				continue
			}

			// 318
			top, err := stack.Top()
			if err != nil {
				log.Fatalf("stack is empty")
			}
			for valK > top && !stack.IsEmpty() && (stack.Len()+(len(s)-k-1) >= 12) {
				stack, _, err = stack.Pop()
				if err != nil {
					// stack is empty, continue
					break
				}
				top, err = stack.Top()
				if err != nil {
					// stack is empty, continue
					break
				}
			}
			if stack.IsEmpty() {
				stack = stack.Push(valK)
				continue
			}

			if stack.Len() >= 12 {
				continue
			}
			// push into stack
			stack = stack.Push(valK)

		}

		// fmt.Println("	Final stack: ", stack)
		temp := float64(0)
		count := 0
		for !stack.IsEmpty() {
			n, err := stack.Top()
			if err != nil {
				log.Fatalln("error stack empty on final pop")
			}
			temp = temp + float64(n)*math.Pow10(count)
			count++
			stack, _, err = stack.Pop()
			if err != nil {
				log.Fatalln("error stack empty on final pop")
			}
		}

		ans += int64(temp)
	}

	fmt.Println("Final ans: ", ans)

}
