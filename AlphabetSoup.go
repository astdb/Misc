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
			testCases, err = strconv.Atoi(inputFileScanner.Text())
			if err != nil {
				log.Fatal(err)

			}

			lineNo++
			continue
		} else {
			// read test case
			testCase := strings.TrimSpace(inputFileScanner.Text())

			// go through testcase char by char, and count the number of times each letter was seen
			// note: need two C's for every single HACKERCUP, so count C's in doubles.
			letterCounts := []int{0, 0, 0, 0, 0, 0, 0, 0}
			subC := 0

			for _, ch := range testCase {
				if string(ch) == "H" {
					letterCounts[0]++
				}

				if string(ch) == "A" {
					letterCounts[1]++
				}

				if string(ch) == "C" {
					subC++

					if subC == 2 {
						letterCounts[2]++
						subC = 0
					}

				}

				if string(ch) == "K" {
					letterCounts[3]++
				}

				if string(ch) == "E" {
					letterCounts[4]++
				}

				if string(ch) == "R" {
					letterCounts[5]++
				}

				if string(ch) == "U" {
					letterCounts[6]++
				}

				if string(ch) == "P" {
					letterCounts[7]++
				}
			}

			// find the least occuring required letter - that would be the number of full HACKERCUPs that can be formed with the letters alread in the soup
			possibleWordFormations := 0

			for k, v := range letterCounts {
				if k == 0 {
					possibleWordFormations = v
				} else {
					if v < possibleWordFormations {
						possibleWordFormations = v
					}
				}
			}

			str.WriteString(fmt.Sprintf("Case #%d: %d\n", lineNo, possibleWordFormations))
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
