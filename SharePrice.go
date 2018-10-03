package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// input filename should be provided on command line
	if len(os.Args) <= 1 {
		fmt.Println("Usage: $> go run SharePrice.go <inputfile>")
		return
	}

	inputFile := os.Args[1] // read input filename from command line

	// open input file and set up scanner to read values
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input file: %v\n", err)
		return
	}
	defer file.Close() // (defer would ensure file handle closure upon program/function return)

	scanner := bufio.NewScanner(file)

	// placeholders for total loss/gain computed for the values from input file
	totalLoss := 0.0
	totalGain := 0.0

	// placeholder for previous line's value, to be used while iterating through the file to compute gain/loss
	// initialized to -1 on assumption that every valid share price is nonnegative
	prevVal := -1.0

	// read through file by line
	for scanner.Scan() {
		// read current line and transform into a float value
		line := strings.TrimSpace(scanner.Text())
		lineVal, err := strconv.ParseFloat(line, 64)

		if err != nil {
			continue // error parsing value, possibly malformed/invalid: move onto next value
		}

		if lineVal < 0 {
			continue // program assumes share price is nonnegative: move onto next value
		}

		if prevVal > -1 {
			// valid previous value, check loss or gain

			if lineVal > prevVal {
				// gain
				totalGain += (lineVal - prevVal)
			}

			if lineVal < prevVal {
				// loss
				totalLoss += (prevVal - lineVal)
			}

			prevVal = lineVal
		} else {
			// first valid value read, initialize
			prevVal = lineVal
		}
	}

	// print gain/loss to stdout
	fmt.Printf("%.2f\n", totalGain)
	fmt.Printf("-%.2f\n", totalLoss)
}
