/*
You own a Goal Parser that can interpret a string command. The command consists of an alphabet of "G", "()" and/or "(al)" in some order. The Goal Parser will interpret "G" as the string "G", "()" as the string "o", and "(al)" as the string "al". The interpreted strings are then concatenated in the original order.

Given the string command, return the Goal Parser's interpretation of command.



Example 1:

Input: command = "G()(al)"
Output: "Goal"
Explanation: The Goal Parser interprets the command as follows:
G -> G
() -> o
(al) -> al
The final concatenated result is "Goal".
Example 2:

Input: command = "G()()()()(al)"
Output: "Gooooal"
Example 3:

Input: command = "(al)G(al)()()G"
Output: "alGalooG"


Constraints:

1 <= command.length <= 100
command consists of "G", "()", and/or "(al)" in some order.
*/

package main

import (
	"log"
	"strings"
)

func main() {
	tests := []string{"G()(al)", "G()()()()(al)", "(al)G(al)()()G"}

	for _, test := range tests {
		log.Printf("interpret(%s) = %s\n", test, interpret(test))
	}
}

func interpret(command string) string {
	var res strings.Builder

	// i := 0
	commandRunes := []rune(command)

	for i := 0; i < len(commandRunes); i++ {
		if commandRunes[i] == 'G' {
			res.WriteString("G")
		} else if i+1 < len(commandRunes) {
			if string(commandRunes[i:i+2]) == "()" {
				res.WriteString("o")
				i++
			}
		} else if i+3 < len(commandRunes) {
			if string(commandRunes[i:i+4]) == "(al)" {
				res.WriteString("(al)")
				i += 3
			}
		}
	}

	return res.String()
}
