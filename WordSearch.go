// given a text corpus, write code to find the occurence frequency of any given word. how would the design change if the code runs many times?

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	// read in the text corpus (probably from a file)
	data, err := ioutil.ReadFile("big.txt")
	check(err)
	book := string(data)

	// split text read into words
	words := strings.Fields(book)
	fmt.Printf("Searching a corpus of %d words..\n", len(words))

	// create map of looked up words to cache results
	results_cache := make(map[string]int)

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter word to search (Ctrl+c to exit): ")
		searchTerm, _ := reader.ReadString('\n')
		searchTerm = strings.TrimSpace(searchTerm)

		// check cache for word
		freq, exists := results_cache[searchTerm]

		if !exists {
			// if not searched for this word's frequency yet, iterate over the words
            fmt.Println("Not found in cache, searching..")
			for _, word := range words {
				if strings.Compare(searchTerm, word) == 0 {
					freq++
				}
			}

            // add to cache
            results_cache[searchTerm] = freq
		} else {
            fmt.Println("Found in cache, returning..")
        }

		fmt.Printf("<%s> was found %d times!\n", searchTerm, freq)
	}
}

func check(e error) {
	if e != nil {
		log.Print(e)
	}
}
