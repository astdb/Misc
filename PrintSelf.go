
// Go program to print its own source code ( inspirted by TPOP (https://www.amazon.com/Practice-Programming-Addison-Wesley-Professional-Computing/dp/020161586X) 9-15)

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"regexp"
)

func main() {
	// ensure source filename entered on command line
	if len(os.Args) < 1 {
		fmt.Println("Cannot read source file name from input arguments")
		return
	}

	// extract source file name from os.Args[0]
	// note: os.Args[0] would either be /log/path/name/progname (go run) or ./progname (compiled)
	// source file name is progname.go
 	r1, _ := regexp.Compile("(?i)[A-Z]+(?-i)$")
	fn := strings.TrimSpace(r1.FindString(os.Args[0])) + ".go"

	// open source file fopr reading
	file, err := os.Open(fn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading self source file: %v\n", err)
		return
	}
	defer file.Close()

	// setup scanner to read source file
	scanner := bufio.NewScanner(file)

	// read each line and print it out
	for scanner.Scan() {
		// read line
		// line := strings.TrimSpace(scanner.Text())	// do not trimspace if output has to have indenting
		line := scanner.Text()
		fmt.Printf("%s\n", line)
	}
}
