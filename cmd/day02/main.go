package main

import (
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/Its-Maniaco/advent-of-code-2025/internal/day02"
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
	fileLoc := "./cmd/day02/"
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
	lines := strings.Split(strings.TrimSpace(data), ",")

	nR := utils.ConvertStringRange(lines)
	merged := utils.UnifyRange(nR)
	fmt.Println(merged)

	fmt.Printf("Enter the part do you want to run:\n1 for Part 1\n2 for Part 2\n Leave blank to run both Part 1 and Part 2\n")
	partRun, err := terminalReader.ReadString('\n')
	if err != nil {
		log.Fatalf("Could not read terminal input to run which part: (%s)", err)
		return
	}
	partRun = strings.TrimSpace(partRun)

	switch partRun {
	case "1":
		day02.Part01(merged)
	case "2":
		day02.Part02(merged)
	default:
		day02.Part01(merged)
		day02.Part02(merged)
	}
}
