/*
International Morse Code defines a standard encoding where each letter is mapped to a series of dots and dashes, as follows: "a" maps to ".-", "b" maps to "-...", "c" maps to "-.-.", and so on.

For convenience, the full table for the 26 letters of the English alphabet is given below:

[".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."]
Now, given a list of words, each word can be written as a concatenation of the Morse code of each letter. For example, "cab" can be written as "-.-.-....-", (which is the concatenation "-.-." + "-..." + ".-"). We'll call such a concatenation, the transformation of a word.

Return the number of different transformations among all words we have.

Example:
Input: words = ["gin", "zen", "gig", "msg"]
Output: 2
Explanation:
The transformation of each word is:
"gin" -> "--...-."
"zen" -> "--...-."
"gig" -> "--...--."
"msg" -> "--...--."

There are 2 different transformations, "--...-." and "--...--.".
*/

package main

import (
	"fmt"
	"os"
)

func main() {
	// read input words list
	if len(os.Args) < 2 {
		fmt.Println("Usage: $> go run UniqueMorseCodeWords.go <word1> <word2> ...")
		return
	}

	inputWords := os.Args[1:]
	fmt.Println(uniqueMorseRepresentations(inputWords))
}

func uniqueMorseRepresentations(words []string) int {
	morseCodes := map[rune]string{
		'a': ".-",
		'b': "-...",
		'c': "-.-.",
		'd': "-..",
		'e': ".",
		'f': "..-.",
		'g': "--.",
		'h': "....",
		'i': "..",
		'j': ".---",
		'k': "-.-",
		'l': ".-..",
		'm': "--",
		'n': "-.",
		'o': "---",
		'p': ".--.",
		'q': "--.-",
		'r': ".-.",
		's': "...",
		't': "-",
		'u': "..-",
		'v': "...-",
		'w': ".--",
		'x': "-..-",
		'y': "-.--",
		'z': "--..",
	}

	codeWordsMap := map[string]int{}

	for i := 0; i < len(words); i++ {
		// get word transformation
		codeWord := ""
		for _, ch := range words[i] {
			codeWord += morseCodes[ch]
		}

		_, codeWordExists := codeWordsMap[codeWord]
		if codeWordExists {
			codeWordsMap[codeWord]++
		} else {
			codeWordsMap[codeWord] = 1
		}
	}

	return len(codeWordsMap)
}
