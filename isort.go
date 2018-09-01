package main

import (
        "fmt"
)

func main() {
        input := []int{5,2,4,6,1,3}
        fmt.Println(isort(input))
}

func isort(a []int) []int {
        if len(a) > 1 {
                for j := 1; j < len(a); j++ {
                        key := a[j]

                        i := j -1
                        for i >= 0 && a[i] > key {
                                a[i+1] = a[i]
                                i--
                        }

                        a[i+1] = key
                }
        }

        return a
}
~     
