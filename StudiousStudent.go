package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"fmt"
)

func main() {
	// read input file name from command line
	if len(os.Args) != 3 {
		log.Fatal("Usage: $>go run StudiousStudent.go input.file output.file")
	}

	inputFileName := os.Args[1]
	outputFileName := os.Args[2]

	// read input file
	inputFile, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	inputFileScanner := bufio.NewScanner(inputFile)
	lineNo := 0
	testCases := 0

	var str strings.Builder
	for inputFileScanner.Scan() {
		inputLine := inputFileScanner.Text()

		if lineNo == 0 {
			// first line of the input file - read number of test cases
			testCases, err = strconv.Atoi(inputLine)
			if err != nil {
				log.Printf("Error reading no of testcases: %v\n", err)
				return
			}

		} else {
			inputLineComponents := strings.Split(inputLine, " ")

			if len(inputLineComponents) > 0 {
				wordCount, err := strconv.Atoi(strings.TrimSpace(inputLineComponents[0]))
				if err != nil {
					log.Printf("Error reading testcase count for testcase #: %d\n", lineNo)
					continue
				}

				if len(inputLineComponents) >= wordCount+1 {
					// well formed testcase
					wordsList := inputLineComponents[1:]
					sort.Strings(wordsList)
					str.WriteString(fmt.Sprintf("Case #%d: %s\n", lineNo, strings.Join(wordsList[:], "")))
				}
			}

		}

		lineNo++
		if lineNo > testCases {
			break
		}

	}

	// write output file
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		log.Fatal(err)
	}

	defer outputFile.Close()

	_, err = outputFile.WriteString(str.String())
	if err != nil {
		log.Fatal(err)
	}

	outputFile.Sync()
}
