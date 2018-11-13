/*
A sentence S is given, composed of words separated by spaces. Each word consists of lowercase and uppercase letters only.

We would like to convert the sentence to "Goat Latin" (a made-up language similar to Pig Latin.)

The rules of Goat Latin are as follows:

If a word begins with a vowel (a, e, i, o, or u), append "ma" to the end of the word.
For example, the word 'apple' becomes 'applema'.

If a word begins with a consonant (i.e. not a vowel), remove the first letter and append it to the end, then add "ma".
For example, the word "goat" becomes "oatgma".

Add one letter 'a' to the end of each word per its word index in the sentence, starting with 1.
For example, the first word gets "a" added to the end, the second word gets "aa" added to the end and so on.
Return the final sentence representing the conversion from S to Goat Latin.



Example 1:

Input: "I speak Goat Latin"
Output: "Imaa peaksmaaa oatGmaaaa atinLmaaaaa"
Example 2:

Input: "The quick brown fox jumped over the lazy dog"
Output: "heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa"


Notes:

S contains only uppercase, lowercase and spaces. Exactly one space between each word.
1 <= S.length <= 150.

*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	testCases := [][]string{{"I speak Goat Latin", "Imaa peaksmaaa oatGmaaaa atinLmaaaaa"}, {"The quick brown fox jumped over the lazy dog", "heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa"}}

	for _, test := range testCases {
		if toGoatLatin(test[0]) == test[1] {
			fmt.Println("Pass")
		} else {
			fmt.Println("Fail", "(Result: \"", toGoatLatin(test[0]), "\")")
		}
	}
}

func toGoatLatin(str string) string {
	// str_runes := []rune(str)
	result := []rune{}
	str_words := strings.Split(str, " ")

	wordIndex := 1
	for _, word := range str_words {
		word_runes := []rune(word)

		if len(word_runes) > 0 {

			if isVowel(word_runes[0]) {
				// if first letter of word is vowel
				suffix := []rune{'m', 'a'}
				word_runes = append(word_runes, suffix...)

			} else {
				// if first letter of word is consonant
				if len(word_runes) > 1 {
					word_start := word_runes[0]
					word_runes = word_runes[1:]
					suffix := []rune{word_start, 'm', 'a'}
					word_runes = append(word_runes, suffix...)
				} else {
					suffix := []rune{'m', 'a'}
					word_runes = append(word_runes, suffix...)
				}

			}

			for i := 0; i < wordIndex; i++ {
				word_runes = append(word_runes, 'a')
			}

			if wordIndex > 1 {
				result = append(result, ' ')
				result = append(result, word_runes...)
			} else {
				result = append(result, word_runes...)
			}

			wordIndex++
		}
	}

	return string(result)
}

func isVowel(ch rune) bool {
	if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' || ch == 'A' || ch == 'E' || ch == 'I' || ch == 'O' || ch == 'U' {
		return true
	}

	return false
}
