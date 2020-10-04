/*
Write a function that takes a string as input and reverse only the vowels of a string.

Example 1:

Input: "hello"
Output: "holle"
Example 2:

Input: "leetcode"
Output: "leotcede"
Note:
The vowels does not include the letter "y".

*/

package main

import (
  "log"
)

func main() {
  tests := []string{"hello", "leetcode"}
  for _, test := range tests {
    // log.Printf("reverseVowels(%s) = %s\n", test, reverseVowels(test))
    log.Printf("reverseVowels(%s) = %s\n", test, reverseVowels2(test))
  }
}

func reverseVowels2(s string) string {
  sRunes := []rune(s)
  vowelIndexList := []int{}

  for i := 0; i < len(sRunes); i++ {
    if isVowel(sRunes[i]) {
      vowelIndexList = append(vowelIndexList, i)
    }
  }

  for i := 0; i < len(vowelIndexList)/2; i++ {
    sRunes[vowelIndexList[i]], sRunes[vowelIndexList[len(vowelIndexList)- i - 1]] = sRunes[vowelIndexList[len(vowelIndexList)- i - 1]], sRunes[vowelIndexList[i]]
  }

  return string(sRunes)
}

func reverseVowels(s string) string {
  // transform input string into unicode char list for easier manipulation
  sRunes := []rune(s)

  // make lists of vowels and their indices from start
  vowelsList  := []rune{}
  vowelIndexList := []int{}

  for i := 0; i < len(sRunes); i++ {
    if isVowel(sRunes[i]) {
      vowelsList = append(vowelsList, sRunes[i])
      vowelIndexList = append(vowelIndexList, i)
    }
  }

  // log.Printf("%v\n", vowelsList)
  // log.Printf("%v\n", vowelIndexList)

  j := 0
  for i := len(vowelIndexList)-1; i >= 0; i-- {
    sRunes[vowelIndexList[i]] = vowelsList[j]
    j++
  }

  return string(sRunes)
}

func isVowel(ch rune) bool {
  if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' || ch == 'A' || ch == 'E' || ch == 'I' || ch == 'O' || ch == 'U' {
    return true
  }

  return false
}
