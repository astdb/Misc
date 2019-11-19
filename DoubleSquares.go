/*
A double-square number is an integer X which can be expressed as the sum of two perfect squares. For example, 10 is a double-square because 10 = 32 + 12. Your task in this problem is, given X, determine the number of ways in which it can be written as the sum of two squares. For example, 10 can only be written as 32 + 12 (we don't count 12 + 32 as being different). On the other hand, 25 can be written as 52 + 02 or as 42 + 32.

Input
You should first read an integer N, the number of test cases. The next N lines will contain N values of X.
Constraints
0 ≤ X ≤ 2147483647
1 ≤ N ≤ 100
Output
For each value of X, you should output the number of ways to write X as the sum of two squares. 
*/

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
		log.Fatal("Usage: $>go run DoubleSquares.go input.file output.file")
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
		num, err := strconv.Atoi(inputFileScanner.Text())
		if err != nil {
			log.Printf("Error reading line %d: %v\n", lineNo, err)
			continue
		}

		if lineNo == 0 {
			testCases = num
			lineNo++
			continue
		}

		if lineNo > testCases {
			break
		}

		y := 0
		squareFactors := [][]int{}
		for {
			xsq := num - (y * y)
			if xsq < 0 {
				break
			}

			// determine if xsq is a full square.
			// if yes, add its square root and y to squareFactors (if not already there)
			x := -1
			for i := 0; i*i <= xsq; i++ {
				if xsq == i*i {
					x = i
					break
				}
			}

			if x != -1 {
				// xsq is a full square
				xyFound := false
				for _, factorSet := range squareFactors {
					if len(factorSet) >= 2 {
						if (x == factorSet[0] || x == factorSet[1]) && (y == factorSet[0] || y == factorSet[1]) {
							xyFound = true
							break
						}
					}
				}

				if !xyFound {
					squareFactors = append(squareFactors, []int{x, y})
				}
			}

			y++
		}

		str.WriteString(fmt.Sprintf("Case #%d: %d\n", lineNo, len(squareFactors)))
		lineNo++
	}

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
