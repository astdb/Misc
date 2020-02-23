/*
Given a string text, you want to use the characters of text to form as many instances of the word "balloon" as possible.

You can use each character in text at most once. Return the maximum number of instances that can be formed.



Example 1:



Input: text = "nlaebolko"
Output: 1
Example 2:



Input: text = "loonbalxballpoon"
Output: 2
Example 3:

Input: text = "leetcode"
Output: 0


Constraints:

1 <= text.length <= 10^4
text consists of lower case English letters only.

*/

package main

import (
	"log"
)

func main() {
	tests := []string{"nlaebolko", "loonbalxballpoon", "leetcode"}

	for k, test := range tests {
		log.Printf("Test %d: maxNumberOfBalloons(%s) == %d\n", k, test, maxNumberOfBalloons(test))
	}
}

func maxNumberOfBalloons(text string) int {
	b_count := 0
	a_count := 0
	l_count := 0
	l_count_sec := 0
	o_count := 0
	o_count_sec := 0
	n_count := 0

	for _, char := range text {
		if char == 'b' {
			b_count++
		}

		if char == 'a' {
			a_count++
		}

		if char == 'l' {
			if l_count_sec == 0 {
				l_count_sec++
			} else {
				l_count_sec = 0
				l_count++
			}
		}

		if char == 'o' {
			if o_count_sec == 0 {
				o_count_sec++
			} else {
				o_count_sec = 0
				o_count++
			}
		}

		if char == 'n' {
			n_count++
		}
	}

	if b_count > 0 && a_count > 0 && l_count > 0 && o_count > 0 && n_count > 0 {
		x := []int{}
		x = append(x, b_count)
		x = append(x, a_count)
		x = append(x, l_count)
		x = append(x, o_count)
		x = append(x, n_count)

		var min int
		for k, i := range x {
			if k == 0 {
				min = i
			} else {
				if i < min {
					min = i
				}
			}
		}

		return min

	} else {
		return 0
	}
}
