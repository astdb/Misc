/*
A bus has n stops numbered from 0 to n - 1 that form a circle. We know the distance between all pairs of neighboring stops where distance[i] is the distance between the stops number i and (i + 1) % n.

The bus goes along both directions i.e. clockwise and counterclockwise.

Return the shortest distance between the given start and destination stops.



Example 1:



Input: distance = [1,2,3,4], start = 0, destination = 1
Output: 1
Explanation: Distance between 0 and 1 is 1 or 9, minimum is 1.


Example 2:



Input: distance = [1,2,3,4], start = 0, destination = 2
Output: 3
Explanation: Distance between 0 and 2 is 3 or 7, minimum is 3.


Example 3:



Input: distance = [1,2,3,4], start = 0, destination = 3
Output: 4
Explanation: Distance between 0 and 3 is 6 or 4, minimum is 4.


Constraints:

1 <= n <= 10^4
distance.length == n
0 <= start, destination < n
0 <= distance[i] <= 10^4

*/

package main

import (
	"log"
)

func main() {
  tests := [][][]int{{{1, 2, 3, 4}, {0}, {1}}, {{1, 2, 3, 4}, {0}, {2}}, {{1, 2, 3, 4}, {0}, {3}}}
  // tests := [][][]int{{{1, 2, 3, 4}, {0}, {1}}}

	for _, test := range tests {
		log.Printf("distanceBetweenBusStops(%v, %d, %d) == %d\n", test[0], test[1][0], test[2][0], distanceBetweenBusStops(test[0], test[1][0], test[2][0]))
	}
}

// this is a new comment
func distanceBetweenBusStops(distance []int, start int, destination int) int {
	n := len(distance)

  // calculate distance one way first..
  log.Printf("Checking clockwise distance..\n")
	dist1 := 0
	curStop := start
  nextStop := getNextStop(start, n)
  log.Printf("Current Stop: %d, Next Stop: %d\n", curStop, nextStop)

	for {
    dist1 += distance[curStop] // dust between start and (start+1)%n
    log.Printf("\tDistance: %d\n", dist1)

		if nextStop == destination {
      log.Printf("\tNextStop (%d) == Destination (%d)\n", nextStop, destination)
			break
		}

		curStop = nextStop
    nextStop = getNextStop(curStop, n)
    log.Printf("Current Stop: %d, Next Stop: %d\n", curStop, nextStop)
	}

	// dist1 += distance[start]  // dust between start and (start+1)%n

  // and now the other..
  log.Printf("\nChecking counter-clockwise distance..\n")
	dist2 := 0
	curStop = start
  prevStop := getPrevStop(curStop, n)
  log.Printf("Current Stop: %d, Previous Stop: %d\n", curStop, prevStop)

	for {
    dist2 += distance[prevStop] // dust between previous stop and current stop
    log.Printf("\tDistance: %d\n", dist2)

		if prevStop == destination {
      log.Printf("\tPreviousStop (%d) == Destination (%d)\n", prevStop, destination)
			break
		}

		curStop = prevStop
    prevStop = getPrevStop(curStop, n)
    log.Printf("Current Stop: %d, Previous Stop: %d\n", curStop, prevStop)
	}

	// return the smaller trip distance
	if dist1 > dist2 {
		return dist2
	}

	return dist1

}

func getPrevStop(stop, n int) int {
	if (stop - 1) >= 0 {
		return (stop - 1)
	}

	return (n - 1)
}

func getNextStop(stop, n int) int {
	return (stop + 1) % n
}
