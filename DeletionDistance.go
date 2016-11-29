
// given two strings, return the minimum number of deletions required to have the same string 

package main

import (
    "fmt"
)

func main() {
    fmt.Println(delete_distance("at", "cat"))
}

func delete_distance(s1 string, s2 string) int {
    s1_r := []rune(s1)
    s2_r := []rune(s2)

    len_s1 := len(s1_r)
    len_s2 := len(s2_r)

    // var d [len_s1+1][len_s2+1]int
    d := make([][]int, len_s1+1)
    for i := range d {
        d[i] = make([]int, len_s2+1)
    }

    for i := 0; i <= len_s1; i++ {
        d[i][0] = i
    }

    for j := 0; j <= len_s2; j++ {
        d[0][j] = j
    }

    for j := 1; j < len_s2; j++ {
        for i := 1; i < len_s1; i++ {
            if s1_r[i] == s2_r[j] {
                d[i][j] = d[i-1][j-1]
            } else {
                d[i][j] = d[i-1][j] + 1
            }
        }
    }

    return d[len_s1][len_s2]
}
