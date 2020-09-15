/*
You have an array of logs.  Each log is a space delimited string of words.

For each log, the first word in each log is an alphanumeric identifier.  Then, either:

Each word after the identifier will consist only of lowercase letters, or;
Each word after the identifier will consist only of digits.
We will call these two varieties of logs letter-logs and digit-logs.  It is guaranteed that each log has at least one word after its identifier.

Reorder the logs so that all of the letter-logs come before any digit-log.  The letter-logs are ordered lexicographically ignoring identifier, with the identifier used in case of ties.  The digit-logs should be put in their original order.

Return the final order of the logs.

 

Example 1:

Input: logs = ["dig1 8 1 5 1","let1 art can","dig2 3 6","let2 own kit dig","let3 art zero"]
Output: ["let1 art can","let3 art zero","let2 own kit dig","dig1 8 1 5 1","dig2 3 6"]
 

Constraints:

0 <= logs.length <= 100
3 <= logs[i].length <= 100
logs[i] is guaranteed to have an identifier, and a word after the identifier.
*/

package main

import (
  "log"
  "strings"
  "strconv"
)

func main() {
  tests := [][]string{{"dig1 8 1 5 1","let1 art can","dig2 3 6","let2 own kit dig","let3 art zero"}}

  for _, test := range tests {
    log.Printf("reorderLogFiles(%v) == %v\n", test, reorderLogFiles(test))
  }
}

func reorderLogFiles(logs []string) []string {
    // each log is a string of tokens.
    // first token is an alphanumeric ID.
    // content after token is either all alphabetic (letter-log) or all numeric (digit-log).
    // order letter logs first, lexicographically ignoring ID (use ID only for tiebreaks). 
    // digit logs must come after all letter logs, in original order.
    digitLogs := []string{}
    letterLogs := []string{}

    for _, thisLog := range logs {
      if isDigitLog(thisLog) {
        digitLogs = append(digitLogs, thisLog)
      } else if isLetterLog(thisLog) {
        letterLogs = append(letterLogs, thisLog)
      }
    }

    log.Printf("reorderLogFiles(): letterLogs: %v\n", letterLogs)
    log.Printf("reorderLogFiles(): digitLogs: %v\n", digitLogs)

    return append(letterLogs, digitLogs...)
}

// indicates if a given log is letter
func isLetterLog(logstr string) bool {
  logstr = strings.TrimSpace(logstr)  // strip any extra whitespace
  if len(logstr) <= 0 {
    return false
  }

  logstrTokens := strings.Split(logstr, " ")  // get tokens
  if len(logstrTokens) <= 1 {
    // each log must have starting ID and at least one token
    return false
  }

  // check if first token after ID can be turned into a number:
  //  - if yes, it must be a digit log (only numbers after ID)
  //  - else, it must be a letter logm (only lowercase letters after ID)
  _, err := strconv.Atoi(logstrTokens[1])
  if err != nil {
    return true
  }

  return false
}

func isDigitLog(logstr string) bool {
  logstr = strings.TrimSpace(logstr)  // strip any extra whitespace
  if len(logstr) <= 0 {
    return false
  }

  logstrTokens := strings.Split(logstr, " ")  // get tokens
  if len(logstrTokens) <= 1 {
    // each log must have starting ID and at least one token
    return false
  }

  // check if first token after ID can be turned into a number:
  //  - if yes, it must be a digit log (only numbers after ID)
  //  - else, it must be a letter log (only lowercase letters after ID)
  _, err := strconv.Atoi(logstrTokens[1])
  if err != nil {
    return false
  }

  return true
}
