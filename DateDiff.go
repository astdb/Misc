package main

import (
  "log"
)

MONTH_DAYS := map[int]int {
  1: 31,
  2: 28,
  3: 31,
  4: 30,
  5: 31,
  6: 30,
  7: 31,
  8: 31,
  9: 30,
  10: 31,
  11: 30,
  12: 31,
}

func main() {

}

// https://stackoverflow.com/questions/25816616/claculate-the-number-of-days-between-two-dates-without-libraries
func diff(startDay, startMonth, startYear, endDay, endMonth, endYear int) int {
  res := (startYear-endYear) * 365
  // res += 
}

func dateDiff(d1, m1, y1, d2, m2, y2 int) {
  // TODO: check if second date is after first date

  for i := y1; i <= y2; i++ {
    if i == y1 {
      // calculate days remaining in y1
      if y1 == y2 && m1 == m2 {
         // dates within the same month of same year
         return d2-d1
      } else if y1 == y2 {
        // dates within same year
        res := 0
        for j := m1; j <= m2; j++ {
          if j == m1 {
            // days until end of start month
            res += days(m1)-d1
          } else if j != m2 {
            res += days(j)
          } else {
            // j == m2
            res += d2
          }
        }

        return res
      } else {
        // dates within different years
         
      }
    }
  }
}

// return number of days in a given month (1-12) and year
func days(m, y) int {
  // TODO: check input

  var days int

  // handle feb
  if m == 2 {
    if leap(y) {
      // leap year
      return 29
    }

    return MONTH_DAYS[m]
  }

  
}

// check if given year is a leap year
func leap(y int) bool {

  return false
}

