/*
Arrays.asList("A", "B", "A", "C", "D", "B", "A");
String findWinner(List<String> votes)

Given a list of strings containing votes cast for a set of candidates, find the winning candidate.

The votes data structure consists of candidate names represented as strings (for simplicity, consider candidate names
to be case-sensitive. i.e. "A" and "a" would be two distinct candaidate names.).

P.S: how could ties be handled?

*/

package main

import (
	"log"
)

func main() {
	// log.Printf("%s\n", findWinner([]string{"A", "B", "A", "C", "D", "B", "A"}))

	tests := [][]string{{"A", "B", "A", "C", "D", "B", "A"}, {"A"}, {"A", "B"}, {"B", "A", "B"}}

	for _, test := range tests {
		log.Printf("findWinner(%s) = %s\n", test, findWinner(test))
	}
}

// findWinner accepts a slice of strings consisting cast votes, and returns the winning candidate name.
func findWinner(votes []string) string {
	// declare map mapping each candidate name to the number of votes they received
	candidateVotes := map[string]int{}

	// iterate through the votes slice and update candidate/vote map
	// keep track of the candidate with most votes (count and name)
	var winnerVotes int
	var winnerName string

	for _, candidate := range votes {
		candidateVotes[candidate]++

		if candidateVotes[candidate] > winnerVotes {
			winnerVotes = candidateVotes[candidate]
			winnerName = candidate
		}
	}

	return winnerName
}
