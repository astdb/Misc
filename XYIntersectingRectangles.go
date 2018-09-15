// given two XY-aligned rectangles, test if they have a non-empty intersection,
// and if so, return the rectangle of the intersection. 

package main

import (
	"fmt"
)

func main() {
	
}

type XYRect struct {
	LineX Line	// line on X-axis
	LineY Line	// line on Y-axis
}

type Line struct {
	X1 int // starting point
	X2 int // ending point
}

// given two line segments on an X or Y coordinate axis, determine if they 
// have a nonempty intersection
func Intersect(l1, l2 Line) bool {
	// either line 1's start is between line 2's start and end
	// or vice-versa
	if (l1.X1 > l2.X1 && l1.X1 < l2.X2) || (l2.X1 > l1.X1 && l2.X1 < l1.X2) {
		return true
	}

return false
}


fmt RecInt(r1, r1 Rectangle) {
	// for two XY-aligned rectangles to intersect, the line segments they cast
	// on X and Y axes must both intersect
	if Intersect(r1.LineX && r2.LineX) && Intersect(r1.LineY && r2.LineY) {
		return true
	}

	return false
}
