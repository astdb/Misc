
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// read input file name from command line
	if len(os.Args) != 3 {
		log.Fatal("Usage: $>go run AlphabetSoup.go input.file output.file")
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
		if lineNo == 0 {
			// read testcase count
			testCases, err := strconv.Atoi(inputFileScanner.Text())
			if err != nil {
				log.Fatal(err)
				
			}
			
			str.WriteString(fmt.Sprintf("%d\n", testCases))
			lineNo++
			continue
		} else {
			// read test case
			testCase := strings.TrimSpace(inputFileScanner.Text())

			str.WriteString(fmt.Sprintf("%s\n", testCase))
		}

		if lineNo > testCases {
			break
		}

		lineNo++
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
