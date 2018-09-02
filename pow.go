// implement x ^ y (double x, int y. may ignore overflow)

package main

import (
    "fmt"
)

func main() {
    fmt.Println(pow(2.0, 30))
    fmt.Println(pow2(2.0, 30))
}

// EPIJ solution
func pow2(x float64, y int) float64 {
	var result float64 = 1.0
	power = y

	if y < 0 {
		power = -power
		x = 1.0 / x
	}

	for power != 0 {
		if (power & 1) != 0 {
			result *= x
		}

		x *= x
		power >>= 1
	}
	
	return result
}

func pow(x float64, y int) float64 {
    if y == 0 {
        return 1
    }

    var res float64
    res = 1.0
    lim := y
    if y < 0 {
        lim = (-1)*lim
    }

    for i := 0; i < lim; i++ {
        res *= x
    }

    if y < 0 {
        fmt.Println(res)
        return (1.0/res)
    }

    return res
}


