// function to return an interleaved version of two given strings

package main

import (
	"fmt"
)

func main() {
	fmt.Println(interleave("123", "abc"))
	fmt.Println(interleave("4567", "d"))
}

func interleave(str1 string, str2 string) string {
	str1_r := []rune(str1)
	str2_r := []rune(str2)
	var str3 []rune
	flag := true

	str1_len := len(str1_r)
	str2_len := len(str2_r)
	i, j := 0, 0

	for {
		if flag == true {
			if i < str1_len {
				str3 = append(str3, str1_r[i])
				i++
			}
			flag = false
		} else {
			if j < str2_len {
				str3 = append(str3, str2_r[j])
				j++
			}
			flag = true
		}

		if i >= str1_len && j >= str2_len {
			break
		}
	}

	return string(str3)
}

//!-
