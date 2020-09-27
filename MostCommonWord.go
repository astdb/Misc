/*
Given a paragraph and a list of banned words, return the most frequent word that is not in the list of banned words.  It is guaranteed there is at least one word that isn't banned, and that the answer is unique.

Words in the list of banned words are given in lowercase, and free of punctuation.  Words in the paragraph are not case sensitive.  The answer is in lowercase. 

Example:

Input: 
paragraph = "Bob hit a ball, the hit BALL flew far after it was hit."
banned = ["hit"]
Output: "ball"
Explanation: 
"hit" occurs 3 times, but it is a banned word.
"ball" occurs twice (and no other word does), so it is the most frequent non-banned word in the paragraph. 
Note that words in the paragraph are not case sensitive,
that punctuation is ignored (even if adjacent to words, such as "ball,"), 
and that "hit" isn't the answer even though it occurs more because it is banned.
 

Note:

1 <= paragraph.length <= 1000.
0 <= banned.length <= 100.
1 <= banned[i].length <= 10.
The answer is unique, and written in lowercase (even if its occurrences in paragraph may have uppercase symbols, and even if it is a proper noun.)
paragraph only consists of letters, spaces, or the punctuation symbols !?',;.
There are no hyphens or hyphenated words.
Words only consist of letters, never apostrophes or other punctuation symbols.
*/

package main

import (
	"log"
	"strings"
)

func main() {
	tests := [][][]string{ {{"Bob hit a ball, the hit BALL flew far after it was hit."}, {"hit"}} }

	for _, test := range tests {
		log.Printf("mostCommonWord(%s, %v) == %s\n", test[0][0], test[1], mostCommonWord(test[0][0], test[1]))
	}
}

func mostCommonWord(paragraph string, banned []string) string {
		// put banned words in a map for O(1) locating
		bannedWordMap := map[string]int{}

		for _, word := range banned {
			word = strings.ToLower(strings.TrimSpace(word))
			_, ok := bannedWordMap[word]
			if ok {
				bannedWordMap[word]++
			} else {
				bannedWordMap[word] = 1
			}
		}

		// split paragraph into words
		paragraphWords := strings.Split(paragraph, " ")

		wordCounts := map[string]int{}
		for _, word := range paragraphWords {
			word = strings.ToLower(strings.TrimSpace(word))
			_, banned := bannedWordMap[strings.TrimSpace(word)]
			if !banned {
				_, counted := wordCounts[word]
				if !counted {
					wordCounts[word] = 1
				} else {
					wordCounts[word]++
				}
			}
		}

		var mostCommonWordd string
		var mostCommonCount int
		
		// for each word/count pair in wordCounts
		i := 0
		for w, c := range wordCounts {
			if i == 0 {
				mostCommonWordd = w
				mostCommonCount = c
				i++
			} else {
				if c > mostCommonCount {
					mostCommonWordd = w
					mostCommonCount = c
				}
			}
		}

		return mostCommonWordd
}
