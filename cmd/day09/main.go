package main

import (
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/Its-Maniaco/advent-of-code-2025/internal/day09"
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
	fileLoc := "./cmd/day09/"
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

	// convert byte to string
	data := string(dataBytes)
	lines := strings.Split(strings.TrimSpace(data), "\n")
	tiles := make([][2]int, 0, len(lines))
	for _, line := range lines {
		numStr := strings.Split(line, ",")
		x, err := strconv.Atoi(numStr[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		y, err := strconv.Atoi(numStr[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		tiles = append(tiles, [2]int{x, y})
	}

	fmt.Printf("Enter the part do you want to run:\n1 for Part 1\n2 for Part 2\n Leave blank to run both Part 1 and Part 2\n")
	partRun, err := terminalReader.ReadString('\n')
	if err != nil {
		log.Fatalf("Could not read terminal input to run which part: (%s)", err)
		return
	}
	partRun = strings.TrimSpace(partRun)

	switch partRun {
	case "1":
		day09.Part01(tiles)
	case "2":
		day09.Part02()
	default:
		day09.Part01(tiles)
		day09.Part02()
	}
}
