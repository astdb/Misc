/*
Given a string s and a string t, check if s is subsequence of t.

A subsequence of a string is a new string which is formed from the original string by deleting some (can be none) of the characters without disturbing the relative positions of the remaining characters. (ie, "ace" is a subsequence of "abcde" while "aec" is not).

Follow up:
If there are lots of incoming S, say S1, S2, ... , Sk where k >= 1B, and you want to check one by one to see if T has its subsequence. In this scenario, how would you change your code?

Credits:
Special thanks to @pbrother for adding this problem and creating all test cases.

 

Example 1:

Input: s = "abc", t = "ahbgdc"
Output: true
Example 2:

Input: s = "axc", t = "ahbgdc"
Output: false
 

Constraints:

0 <= s.length <= 100
0 <= t.length <= 10^4
Both strings consists only of lowercase characters.

*/

package main

import (
  "log"
)

func main() {
  tests := [][]string{{"abc", "ahbgdc"}, {"axc", "ahbgdc"}}

  for _, test := range tests {
    log.Printf("isSubsequence(%s, %s) == %v\n", test[0], test[1], isSubsequence(test[0], test[1]))
  }
}

func isSubsequence(s string, t string) bool {
  s_runes := []rune(s)
  t_runes := []rune(t)

  curJ := 0
  found := 0
  for i := 0; i < len(s_runes); i++ {
    ch := s_runes[i]

    for j := curJ; j < len(t_runes); j++ {
      if t_runes[j] == s_runes[i] {
        curJ = j + 1
        found++
        break
      }
    }
  }

  if found == len(s_runes) {
    return true
  }

  return false
}
