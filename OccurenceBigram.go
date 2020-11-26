/*
Given words first and second, consider occurrences in some text of the form "first second third", where second comes immediately after first, and third comes immediately after second.

For each such occurrence, add "third" to the answer, and return the answer.



Example 1:

Input: text = "alice is a good girl she is a good student", first = "a", second = "good"
Output: ["girl","student"]
Example 2:

Input: text = "we will we will rock you", first = "we", second = "will"
Output: ["we","rock"]


Note:

1 <= text.length <= 1000
text consists of space separated words, where each word consists of lowercase English letters.
1 <= first.length, second.length <= 10
first and second consist of lowercase English letters.

*/

package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	tests := []*TestObj{&TestObj{Text: "alice is a good girl she is a good student", First: "a", Second: "good"}, &TestObj{Text: "we will we will rock you", First: "we", Second: "will"}}

	for _, test := range tests {
		log.Printf("findOccurences(%s, %s, %s) = %v\n", test.Text, test.First, test.Second, findOcurrences(test.Text, test.First, test.Second))
	}
}

func findOcurrences(text string, first string, second string) []string {
	bigramMap := map[string][]string{}

	textTokens := strings.Split(text, " ")

	for i := 0; i < len(textTokens); i++ {
		if (i + 2) < len(textTokens) {
			bigram := fmt.Sprintf("%s %s", strings.TrimSpace(textTokens[i]), strings.TrimSpace(textTokens[i+1]))

			_, mapped := bigramMap[bigram]
			if mapped {
				if newBigramSuccessor(bigramMap[bigram], textTokens[i+2]) {
					bigramMap[bigram] = append(bigramMap[bigram], textTokens[i+2])
				}
			} else {
				bigramMap[bigram] = []string{strings.TrimSpace(textTokens[i+2])}
			}
		}
	}

	return bigramMap[fmt.Sprintf("%s %s", first, second)]
}

func newBigramSuccessor(tokenList []string, token string) bool {
	for _, tok := range tokenList {
		if strings.TrimSpace(tok) == strings.TrimSpace(token) {
			return false
		}
	}

	return true
}

type TestObj struct {
	Text   string
	First  string
	Second string
}
