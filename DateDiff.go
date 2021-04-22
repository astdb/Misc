package main

import (
	// "log"
)

var MONTH_DAYS = map[int]int{
	1:  31,
	2:  28,
	3:  31,
	4:  30,
	5:  31,
	6:  30,
	7:  31,
	8:  31,
	9:  30,
	10: 31,
	11: 30,
	12: 31,
}

func main() {

}

// https://stackoverflow.com/questions/25816616/claculate-the-number-of-days-between-two-dates-without-libraries
// func diff(startDay, startMonth, startYear, endDay, endMonth, endYear int) int {
// 	res := (startYear - endYear) * 365
// 	// res +=
// }

func dateDiff(d1, m1, y1, d2, m2, y2 int) int {
	// TODO: check if second date is after first date
	res := 0

	for i := y1; i <= y2; i++ {
		// if starting year
		if i == y1 {
			// calculate days remaining in y1
			if y1 == y2 && m1 == m2 {
				// dates within the same month of same year
				res = d2 - d1
				return res

			} else if y1 == y2 {
				// dates within same year, but different months
				// res := 0

				// per month
				for j := m1; j <= m2; j++ {
					if j == m1 {
						// days until end of start month
						res += daysMonth(m1, i) - d1
					} else if j != m2 {
						// middle month(s)
						res += daysMonth(j, i)
					} else {
						// j == m2
						// last month
						res += d2
					}
				}

				return res

			} else {
				// dates within different years - calculate days until end of this year
				// for j := m1; j <= m2; j++ {
				for j := m1; j <= 12; j++ {
					if j == m1 {
						// days until end of start month
						res += daysMonth(m1, i) - d1

					} else if j != m2 {
						res += daysMonth(j, i)

					} else {
						// j == m2
						res += d2
					}
				}
			}
		} else if i != y2 {
			// midde year(s)
			res += daysYear(i)

		} else if i == y2 {
			// calculate days remaining in last year
			for j := 1; j <= m2; j++ {
				if j != m2 {
					// not the last month
					res += daysMonth(j, y2)
				} else if j == m2 {
					res += d2
				}
			}
		}

		// return res
	}

  return res
}

// return number of days in a given month (1-12) and year
func daysMonth(m, y int) int {
	// TODO: check input
	// var MONTH_DAYS = map[int]int{
	// 	1:  31,
	// 	2:  28,
	// 	3:  31,
	// 	4:  30,
	// 	5:  31,
	// 	6:  30,
	// 	7:  31,
	// 	8:  31,
	// 	9:  30,
	// 	10: 31,
	// 	11: 30,
	// 	12: 31,
	// }

	// var days int

	// handle feb
	if m == 2 && leap(y) {
		return 29
	}

  return MONTH_DAYS[m]
}

// return number of days in a given year
func daysYear(y int) int {
	if leap(y) {
		return 366
	}

	return 365
}

// check if given year is a leap year
func leap(y int) bool {
	// return ((year % 4 == 0) && (year % 100 != 0)) || (year % 400 == 0);
	return ((y%4 == 0) && (y%100 != 0)) || (y%400 == 0)

}
