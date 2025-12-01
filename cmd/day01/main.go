package main

import (
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/Its-Maniaco/advent-of-code-2025/internal/day01"
	"github.com/Its-Maniaco/advent-of-code-2025/internal/utils"
)

func main() {
	// take flag for input file or test file
	terminalReader := utils.GetTerminalInputReader()
	fmt.Println("Enter 0 to run on your main input\nEnter 1 (or leave blank) to run on your test input")
	isTestInput, err := terminalReader.ReadString('\n')
	if err != nil {
		log.Fatalf("Could not read terminal input. (%s)", err)
		return
	}

	isTestInput = strings.TrimSpace(isTestInput)

	isTest := true
	fileLoc := "./cmd/day01/"
	if isTestInput != "0" {
		isTest = false
	}

	fmt.Println("isTest T or F: ", isTest)
	fs, err := utils.OpenInput(fileLoc, isTest)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
		return
	}

	// read complete file
	dataBytes, err := io.ReadAll(fs)
	if err != nil {
		log.Fatalf("failed to read file: %s", err)
		return
	}

	// convert byte to string
	data := string(dataBytes)
	lines := strings.Split(strings.TrimSpace(data), "\n")

	num := make([]int, len(lines))
	// change string to int input
	for i, sNum := range lines {
		num[i], err = strconv.Atoi(string(sNum)[1:])
		if err != nil {
			log.Fatalf("Failed to convert string to int: %s", err)
			return
		}
		if string(sNum)[0] == 'L' {
			num[i] *= -1
		}
	}

	day01.Part01(num)

}
