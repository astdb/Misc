/*
Given a valid (IPv4) IP address, return a defanged version of that IP address.

A defanged IP address replaces every period "." with "[.]".



Example 1:

Input: address = "1.1.1.1"
Output: "1[.]1[.]1[.]1"
Example 2:

Input: address = "255.100.50.0"
Output: "255[.]100[.]50[.]0"


Constraints:

The given address is a valid IPv4 address.

*/

package main

import (
	"log"
	"strings"
)

func main() {
	tests := []string{"1.1.1.1", "255.100.50.0"}
	for i, test := range tests {
		log.Printf("Test #%d: defangIPaddr(%s) == %s\n", i, test, defangIPaddr(test))
	}
}

func defangIPaddr(address string) string {
	var str strings.Builder

	for _, ch := range address {
		if ch == '.' {
			str.WriteString("[.]")
		} else {
			str.WriteString(string(ch))
		}
	}

	return str.String()
}
