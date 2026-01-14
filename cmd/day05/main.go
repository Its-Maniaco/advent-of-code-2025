package main

import (
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/Its-Maniaco/advent-of-code-2025/internal/day05"
	"github.com/Its-Maniaco/advent-of-code-2025/internal/utils"
)

func main() {
	terminalReader := utils.GetTerminalInputReader()
	fmt.Println("Enter 0 to run on your main input\nEnter 1 (or leave blank) to run on your test input")
	isTestInput, err := terminalReader.ReadString('\n')
	if err != nil {
		log.Fatalf("Could not read terminal input. (%s)", err)
		return
	}

	isTestInput = strings.TrimSpace(isTestInput)

	isTest := true
	fileLoc := "./cmd/day05/"
	if isTestInput != "0" {
		isTest = false
	}

	fs, err := utils.OpenInput(fileLoc, isTest)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
		return
	}
	defer fs.Close()

	// read complete file
	dataBytes, err := io.ReadAll(fs)
	if err != nil {
		log.Fatalf("failed to read file: %s", err)
		return
	}

	fmt.Printf("Enter the part do you want to run:\n1 for Part 1\n2 for Part 2\n Leave blank to run both Part 1 and Part 2\n")
	partRun, err := terminalReader.ReadString('\n')
	if err != nil {
		log.Fatalf("Could not read terminal input to run which part: (%s)", err)
		return
	}
	partRun = strings.TrimSpace(partRun)

	rangeStr := []string{}
	nums := []int{}
	flag := false
	input := strings.Split(string(dataBytes), "\n")
	for _, v := range input {
		if v == "" {
			flag = true
			continue
		}
		if !flag {
			rangeStr = append(rangeStr, v)
		} else {
			n, err := strconv.Atoi(v)
			if err != nil {
				fmt.Printf("error converting string to num: %v\n", err)
				return
			}

			nums = append(nums, n)
		}
	}

	switch partRun {
	case "1":
		day05.Part01(rangeStr, nums)
	case "2":
		day05.Part02(rangeStr, nums)
	default:
		day05.Part01(rangeStr, nums)
		day05.Part02(rangeStr, nums)
	}
}
