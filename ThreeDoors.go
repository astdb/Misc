// program to check if you actually get better odds by changing doors if shown one to be empty
// backstory: the popular riddle goes as that at a game show you're shown three doors, behind one 
// there's an expensive gift - the others are empty. once you select one, the host opens one of the 
// remaining doors and shows that its empty, and gives you a chance to revise your choice - 
// do you stick with your original selection or do you swap? 

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	totalRuns := 0.0
	totalWins := 0.0
	change := false

	for i := 0; i < 10000; i++ {
		totalRuns++

		// initialize three doors with a gift behind one 
		doors := []int{0, 0, 0}	// doors are slice indexes 0, 1, and 2 - gift will be changed to 1
		giftIndex := rand.Intn(3)
		doors[giftIndex] = 1

		// contestant selects a door 
		contestantChoice := rand.Intn(3)

		// host demonstrates the empty door from the remaining
		empty := demoEmptyDoor(doors, contestantChoice)

		// change or not
		contestantChoice = changeOrKeep(contestantChoice, empty, change)

		// outcome 
		if doors[contestantChoice] == 1 {
			totalWins++
		}
	}

	odds := (totalWins * 100) / totalRuns
	fmt.Printf("Change: %v\nRuns: %.0f\nWins: %.0f\nOdds: %.2f%%\n", change, totalRuns, totalWins, odds)
}

// given contestant's choice, shown empty door, and a flag denoting to change selection or not, return a selection
// (changed if flag set)
func changeOrKeep(choice int, empty int, change bool) int {
	if change {
		return choice
	}

	for i := 0; i < 3; i++ {
		if i == choice || i == empty {
			continue
		}

		return i
	}

	return 10
}

func demoEmptyDoor(doors []int, choice int) int {
	for i := 0; i < 3; i++ {
		if i == choice {
			continue
		}

		if doors[i] == 0 {
			return i
		}
	}

	return 10
}
