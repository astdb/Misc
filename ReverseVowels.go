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
    log.Printf("reverseVowels(%s) = %s\n", test, test)
  }
}

func reverseVowels(s string) string {
   
}

