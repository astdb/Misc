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
		log.Fatal("Usage: $>go run PegGame.go input.file output.file")
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
				log.Fatal(err)
			}

			if testCases < 1 || testCases > 100 {
				log.Fatal("N outside parameters.")
			}

			str.WriteString(fmt.Sprintf("%d\n", testCases))
		} else {
			inputLineComponents := strings.Split(inputLine, " ")
			// r := 0
			// c := 0
			if len(inputLineComponents) >= 4 {
				r, err := strconv.Atoi(strings.TrimSpace(inputLineComponents[0]))
				if err != nil {
					log.Fatal(err)
				}

				c, err := strconv.Atoi(strings.TrimSpace(inputLineComponents[1]))
				if err != nil {
					log.Fatal(err)
				}

				if r < 3 || r > 100 {
					log.Fatal("R outside parameters.")
				}

				if c < 3 || c > 100 {
					log.Fatal("C outside parameters.")
				}

				k, err := strconv.Atoi(strings.TrimSpace(inputLineComponents[2]))
				if err != nil {
					log.Fatal(err)
				}

				m, err := strconv.Atoi(strings.TrimSpace(inputLineComponents[3]))
				if err != nil {
					log.Fatal(err)
				}

				str.WriteString(fmt.Sprintf("%d %d %d %d ", r, c, k, m))

				if len(inputLineComponents) >= (4 + m*2) {
					i := 3
					for j := 0; j < m; j++ {
						str.WriteString(fmt.Sprintf("{%s %s} ", inputLineComponents[i+1], inputLineComponents[i+2]))
						i += 2
					}

					str.WriteString("\n")
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
