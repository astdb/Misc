/*
    Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
    The brackets must close in the correct order, "()" and "()[]{}" are all valid but "(]" and "([)]" are not.
*/

package main

import (
	"fmt"
	"os"
)

func main () {
	// strategy: every time an opening paren is found, push it onto a stack. 
	// everytime a closing paren is found, pop the stack and chck if its the matching opening paren - if not, return false.
	// at the end of parsing the string, an empty stack would show that the string had matching parens. 

	// read input string
	if len(os.Args) < 2 {
		fmt.Println("Usage: $> go run ValidParen.go <parenthesis_string>")
		return
	}

	input := os.Args[1]
	fmt.Printf("%v\n", validParens(input))	
}

func validParens(s string) bool {
	// stack for storing parenthesis
	parenStack := []rune{}

	// iterate through the string
	for _,ch := range s {
		// if an opening paren is found, push onto stack
		if ch == '(' || ch == '{' || ch == '[' {
			parenStack = append(parenStack, ch)
		}

		// if a closing paren is found, pop the stack and check its the correct opening paren
		if ch == ')' {
			if len(parenStack) <= 0 {
				return false
			}

			if parenStack[len(parenStack)-1] != '(' {
				return false
			}

			parenStack = parenStack[0:len(parenStack)-1]
		}

		// if a closing paren is found, pop the stack and check its the correct opening paren
		if ch == '}' {
			if len(parenStack) <= 0 {
				return false
			}

			if parenStack[len(parenStack)-1] != '{' {
				return false
			}

			parenStack = parenStack[0:len(parenStack)-1]
		}

		// if a closing paren is found, pop the stack and check its the correct opening paren
		if ch == ']' {
			if len(parenStack) <= 0 {
				return false
			}

			if parenStack[len(parenStack)-1] != '[' {
				return false
			}

			parenStack = parenStack[0:len(parenStack)-1]
		}
	}

	if len(parenStack) == 0 {
		return true
	}

	return false
}
