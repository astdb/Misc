/*
You are given a string word that consists of digits and lowercase English letters.

You will replace every non-digit character with a space. For example, "a123bc34d8ef34" will become " 123  34 8  34". Notice that you are left with some integers that are separated by at least one space: "123", "34", "8", and "34".

Return the number of different integers after performing the replacement operations on word.

Two integers are considered different if their decimal representations without any leading zeros are different.



Example 1:

Input: word = "a123bc34d8ef34"
Output: 3
Explanation: The three different integers are "123", "34", and "8". Notice that "34" is only counted once.
Example 2:

Input: word = "leet1234code234"
Output: 2
Example 3:

Input: word = "a1b01c001"
Output: 1
Explanation: The three integers "1", "01", and "001" all represent the same integer because
the leading zeros are ignored when comparing their decimal values.


Constraints:

1 <= word.length <= 1000
word consists of digits and lowercase English letters.

*/

package main

import (
	"log"
	"strconv"
)

func main() {
	tests := []string{"", "1", "1a", "a123bc34d8ef34", "leet1234code234", "a1b01c001"}
	for _, test := range tests {
		log.Printf("numDifferentIntegers(\"%s\") = %d\n", test, numDifferentIntegers(test))
	}
}

func numDifferentIntegers(word string) int {
	// iterate over word, character by character
	numCounts := map[int]int{}

	numStr := false
	numStrStart := 0
	numStrEnd := 0

	for i, ch := range word {
		_, err := strconv.Atoi(string(ch))
		if err == nil {
			// ch is a valid numeric character
			if numStr {
				// add to current numeric string

			} else {
				numStr = true
				numStrStart = i
			}
		} else {
			// ch not numeric
			if numStr {
				numStr = false
				numStrEnd = i

				num, err1 := strconv.Atoi(word[numStrStart:numStrEnd])
				if err1 != nil {
					log.Printf("Error converting %s to int.\n", word[numStrStart:numStrEnd])
				}
				_, ok := numCounts[num]
				if ok {
					numCounts[num]++
				} else {
					numCounts[num] = 1
				}

				numStrStart = 0
				numStrEnd = 0
			}
		}
	}

	if numStr {
		num, err1 := strconv.Atoi(word[numStrStart:])
		if err1 != nil {
			log.Printf("Error converting %s to int.\n", word[numStrStart:])
		}
		_, ok := numCounts[num]
		if ok {
			numCounts[num]++
		} else {
			numCounts[num] = 1
		}

	}

	// log.Println(numCounts)
	return len(numCounts)
}
