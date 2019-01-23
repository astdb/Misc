/*
We are to write the letters of a given string S, from left to right into lines. Each line has maximum width 100 units, and if writing a letter would cause the width of the line to exceed 100 units, it is written on the next line. We are given an array widths, an array where widths[0] is the width of 'a', widths[1] is the width of 'b', ..., and widths[25] is the width of 'z'.

Now answer two questions: how many lines have at least one character from S, and what is the width used by the last such line? Return your answer as an integer list of length 2.



Example :
Input:
widths = [10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10]
S = "abcdefghijklmnopqrstuvwxyz"
Output: [3, 60]
Explanation:
All letters have the same length of 10. To write all 26 letters,
we need two full lines and one line with 60 units.
Example :
Input:
widths = [4,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10]
S = "bbbcccdddaaa"
Output: [2, 4]
Explanation:
All letters except 'a' have the same length of 10, and
"bbbcccdddaa" will cover 9 * 10 + 2 * 4 = 98 units.
For the last 'a', it is written on the second line because
there is only 2 units left in the first line.
So the answer is 2 lines, plus 4 units in the second line.


Note:

The length of S will be in the range [1, 1000].
S will only contain lowercase letters.
widths is an array of length 26.
widths[i] will be in the range of [2, 10].
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(numberOfLines([]int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, "abcdefghijklmnopqrstuvwxyz"))
	fmt.Println(numberOfLines([]int{4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, "bbbcccdddaaa"))
}

func numberOfLines(widths []int, str string) []int {
	// store letter widths in a letter-indexed hashmap
	if len(widths) < 26 {
		return nil
	}
	
	widthsMap := map[rune]int{}
	widthsMap['a'] = widths[0]
	widthsMap['b'] = widths[1]
	widthsMap['c'] = widths[2]
	widthsMap['d'] = widths[3]
	widthsMap['e'] = widths[4]
	widthsMap['f'] = widths[5]
	widthsMap['g'] = widths[6]
	widthsMap['h'] = widths[7]
	widthsMap['i'] = widths[8]
	widthsMap['j'] = widths[9]
	widthsMap['k'] = widths[10]
	widthsMap['l'] = widths[11]
	widthsMap['m'] = widths[12]
	widthsMap['n'] = widths[13]
	widthsMap['o'] = widths[14]
	widthsMap['p'] = widths[15]
	widthsMap['q'] = widths[16]
	widthsMap['r'] = widths[17]
	widthsMap['s'] = widths[18]
	widthsMap['t'] = widths[19]
	widthsMap['u'] = widths[20]
	widthsMap['v'] = widths[21]
	widthsMap['w'] = widths[22]
	widthsMap['x'] = widths[23]
	widthsMap['y'] = widths[24]
	widthsMap['z'] = widths[25]

	curLineRem := 100   // characters remaining on the current being written
	lastLineLength := 0 // length of the last line written
	lineCount := 1
	for _, ch := range str {
		// insert character to current line if there's space
		if curLineRem-widthsMap[ch] >= 0 {
			curLineRem -= widthsMap[ch]
			lastLineLength += widthsMap[ch]
		} else {
			// start new line and insert character
			curLineRem = 100
			lastLineLength = 0
			lineCount += 1

			curLineRem -= widthsMap[ch]
			lastLineLength += widthsMap[ch]
		}

	}

	result := []int{lineCount, lastLineLength}
	return result

}
