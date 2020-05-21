/*
Given a string date representing a Gregorian calendar date formatted as YYYY-MM-DD, return the day number of the year.



Example 1:

Input: date = "2019-01-09"
Output: 9
Explanation: Given date is the 9th day of the year in 2019.
Example 2:

Input: date = "2019-02-10"
Output: 41
Example 3:

Input: date = "2003-03-01"
Output: 60
Example 4:

Input: date = "2004-03-01"
Output: 61


Constraints:

date.length == 10
date[4] == date[7] == '-', and all other date[i]'s are digits
date represents a calendar date between Jan 1st, 1900 and Dec 31, 2019.
*/

package main

import (
	"log"
	"strconv"
	"strings"
)

func main() {
	tests := []string{"2019-01-09", "2019-02-10", "2003-03-01", "2004-03-01"}

	for _, test := range tests {
		log.Printf("dayOfYear(%s) == %d\n", test, dayOfYear(test))
	}
}

func dayOfYear(date string) int {
	dateComps := strings.Split(date, "-")

	if len(dateComps) == 3 {
		dateYear, err := strconv.Atoi(strings.TrimSpace(dateComps[0]))
		if err != nil {
			log.Fatal(err)
		}

		monthDays := map[int]int{}
		monthDays[1] = 31

		if leapYear(dateYear) {
			monthDays[2] = 29
		} else {
			monthDays[2] = 28
		}
		monthDays[3] = 31
		monthDays[4] = 30
                monthDays[5] = 31
                monthDays[6] = 30
		monthDays[7] = 31
		monthDays[8] = 31
		monthDays[9] = 30
		monthDays[10] = 31
		monthDays[11] = 30
		monthDays[12] = 31

		dateMonth, err := strconv.Atoi(strings.TrimSpace(dateComps[1]))
		if err != nil {
			log.Fatal(err)
		}

		dateDay, err := strconv.Atoi(strings.TrimSpace(dateComps[2]))
		if err != nil {
			log.Fatal(err)
		}

		dayNo := dateDay

		for i := 1; i < dateMonth; i++ {
			dayNo += monthDays[i]
		}

                return dayNo
	}

	return -1
}

func leapYear(year int) bool {
	if year%400 == 0 {
		return true
	}

	if year%100 == 0 {
		return false
	}

	if year%4 == 0 {
		return true
	}

	return false
}
