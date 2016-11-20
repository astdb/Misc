// function to calculate the nearest power to two to a given integer

package main

import (
    "fmt"
    "math"
)

func main(){
    fmt.Println(nearest_power_of_two(10))
    fmt.Println(nearest_power_of_two(2))
    fmt.Println(nearest_power_of_two(100))
}

func nearest_power_of_two(n int) int {
    i, res := 0, 0

    for {
        powVal := int(math.Pow(2, float64(i)))

        if powVal > n {
            return res
        }

        res = powVal
        i++
    }
}

//!-
