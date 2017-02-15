/*
    Write a function to swap a number in-place (i.e. without a temporary variable)
*/

package main

import (
    "fmt"
)

func main(){
    j, k := 1, 4
    fmt.Printf("Pre-swap: j = %d, k = %d\n", j, k)
    // InPlaceSwapXOR(&j, &k)
    InPlaceSwapARITH(&k, &j)
    fmt.Printf("Pre-swap: j = %d, k = %d\n", j, k)
}

func InPlaceSwapXOR(i1,i2 *int) {
    fmt.Println("Swapping with InPlaceSwapXOR")
    *i1 = *i1 ^ *i2
    *i2 = *i2 ^ *i1
    *i1 = *i1 ^ *i2
}

func InPlaceSwapARITH(i1,i2 *int) {
    // Note: i1 has to be the larger value
    fmt.Println("Swapping with InPlaceSwapARITH")
    *i1 = *i1 - *i2
    *i2 = *i1 + *i2
    *i1 = *i2 - *i1
}
