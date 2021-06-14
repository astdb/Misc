/*
Imagine we are building an application that is used by many clients. We want to avoid any one client being able to
overload the system by sending too many requests, so we enforce a per-customer rate limit. The rate limit is defined as:

“Each customer can make X requests per Y seconds”

Assuming that customer ID is extracted somehow from the request, implement the following function.


boolean rateLimit(int customerId)


SOLUTION
--------

The rate limit history can be stored in the below data structure:

var requestRates map[int]map[int]int{}

This maps a client ID to another map storing unix second to requests made within that second.

When a request comes in, check if the client ID is present in requestRates. If so, retrieve the
second -> request map for it; if not, insert client ID with current second mapping to a req count of one.

Check if the second in the second map is the current second. If yes, increment it, if not add current
second to it with a request count of one. Check if req count exceeds the rate limit, and return appropriate value.

This could result in the inner map structure mapped to by the client ID key growing to thousands of entries
for busy periods. This could be replaced by the below simpler requestRates structure.

var requestRates map[int][]int{}

In this version, the client ID points to an int slice. This slice only needs to have two elements: one
for then Unix second and another for the number of requests made in that second.

When a request comes in the program would retrieve the second/req slice by client ID. Then it would check if
the first element is the current second and update the seconds count if so. Else it’ll overwrite the first element
with the current second and reset req count to one.

*/

package main

import (
	// "log"
	"time"
)

// REQ_RATES tracks number of API requests made per client ID.
// It maps a client ID to a struct storing a Unix second and requests made within that second.
var REQ_RATES map[int]*SecReq

// number of requests a client allowed to make per second
var RATE_LIMIT int64 = 1000

func main() {

}

func rateLimit(custID int) bool {
	var limit bool

	// check if this client ID is present in REQ_RATES
	secReqs, present := REQ_RATES[custID]

	if !present {
		// customer ID not found in req rates - initialize
		REQ_RATES[custID] = &SecReq{Sec: time.Now().Unix(), Reqs: 1}

		if 1 < RATE_LIMIT {
			// req count below limit - don't rate limit
			limit = false
			return limit
		}

		// req count >= limit - rate limit
		limit = true
		return limit
	} else {
		// custID available in REQ_RATES
		// read current unix sec
		curSec := time.Now().Unix()

		if secReqs != nil {
			if secReqs.Sec == curSec {
				// count available for requests made by this client for current second
				secReqs.Reqs++ // add this req to count

				if secReqs.Reqs < RATE_LIMIT {
					// req count < rate limit - don't throttle
					limit = false
					return limit
				} else {
					// req count over limit per sc - throttle
					limit = true
					return limit
				}
			} else {
				// count for current second not available - initialize req count for current sec
				secReqs.Sec = curSec
				secReqs.Reqs = 1
			}
		} else {
			REQ_RATES[custID] = &SecReq{Sec: curSec, Reqs: 1}

			if 1 < RATE_LIMIT {
				limit = false
				return limit
			}
		}
	}

	return limit
}

// struct type mapping count of requests made in a given UNIX second
type SecReq struct {
	Sec  int64
	Reqs int64
}
