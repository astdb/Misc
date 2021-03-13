/*
Given an integer array nums, return the largest perimeter of a triangle with a non-zero area, formed from three of these lengths. If it is impossible to form any triangle of a non-zero area, return 0.

 

Example 1:

Input: nums = [2,1,2]
Output: 5
Example 2:

Input: nums = [1,2,1]
Output: 0
Example 3:

Input: nums = [3,2,3,4]
Output: 10
Example 4:

Input: nums = [3,6,2,3]
Output: 8
 

Constraints:

3 <= nums.length <= 104
1 <= nums[i] <= 106

*/

package main

import (
  "log"
  "math"
)

func main() {
  tests := [][]int{{2,1,2}, {1,2,1}, {3,2,3,4}, {3,6,2,3}, {741543,812711,204837,321987,64537,878008,221451,327106,519590,632854,258421,777138,648326,163607,970976,160711,832045,458756,21938,914098,835982,430595,342469,265138,225707,494713,495327,112759,147046,115957,385779,852227,452239,508931,825955,438955,71058,49515,379521,239031,124516,197904,520779,309104,220997,17259,834188,589665,971314,723219,738549,348958,18166,494987,445213,202268,65738,575195,142499,720795,857949,241693,964781,288878,355721,149425,569030,128272,250882,766638,781278,954830,202897,499908,242161,836044,882445,149046,973168,308861,695969,397252,694236,786512,225385,758712,575288,903009,777254,312041,90724,997241,199369,901039,976660,433213,47420,768493,3320,666795,996456,436528,122410,47092,648548,57669,390447,297586,355856,968592,522598,278337,961624,893997,738786,548212,48845,344374,666675,86626,613747,825657,259026,279646,976958,973717,576245,18265,834727,889977,196365,712752,654079,575652,921270,333401,212866,952770,876057,898947,338716,540373,494868,650588,500175,405697,511664,852874,599093,852449,542057,237342,63644,599522,855191,795427,134938,206058,649805,567419,629528,605327,929921,258433,135699,148359,469488,580824,645548,132206,61879,306148,739102,286337,29256,883458,684705,952,481629,417700,689973,898502,953980,283946,97225,4568,356564,356320,828541,145931,458524,286262,334333,882861,672029,595393,174280,335137,533160,724925,415784,62706,760414,813218,261695,950372,887523,250470,109081,755534,105815,380556,654806,722973,804792,131818,115923,999825,583355,425454,80047,572825,345627,763985,482663,389383,401038,280261,617918,868561,854346,47395,686237,237301,583331,160550,268421,472463,307708,883420,833792,31213,74926,175036,512415,242086,780627,782536,767326,893268,944117,347183,933039,214984,75597,121411,296192,901229,798467,291665,145536,974275,471494,284938,672862,855188,475079,134455,180062,200699,953848,596392,827969,550705,892445,428671,81371,89978,324386,471471,771795,590293,539513,392215,698312,234095,155651,498085,407382,409059,45130,190111,976204,283988,621650,890320,565746,405086,320482,603277}}

  for _, test := range tests {
    log.Printf("largestPerimeter(%v) = %d\n", test, largestPerimeter(test))
    /* log.Printf("---------------------------Combinations for %v------------------------------------\n", test)
    res := []int{0,0,0}
    sideCombs(test, 3, 0, res)
    log.Printf("%v\n", allSideCombs)
    allSideCombs = [][]int{} */
  }
}

func largestPerimeter(nums []int) int {
  // log.Printf("nums: %v\n", nums)

  // generate all side combinations
  res := []int{0,0,0}
  sideCombs(nums, 3, 0, res)
  // var maxArea float64
  var maxPer int

  // log.Printf("%v\n", allSideCombs)

  for _, sideComb := range allSideCombs {
    area := getArea(sideComb)
    perim := getPerim(sideComb)

    // log.Printf("sides: %v / area: %f, perim: %d\n", sideComb, area, perim)

    if area > 0 {
      if perim > maxPer {
        maxPer = perim
      }
    }
  }

  allSideCombs = [][]int{}
  return maxPer
}

// return triangle area given side lengths
// https://www.mathopenref.com/heronsformula.html
func getArea(sideComb []int) float64 {
  a := float64(sideComb[0])
  b := float64(sideComb[1])
  c := float64(sideComb[2])

  p := (a+b+c)/2

  return math.Sqrt(p*(p-a)*(p-b)*(p-c))
}

func getPerim(sideComb []int) int {
  var perim int

  for _, v := range sideComb {
    perim += v
  }

  return perim
}


var allSideCombs [][]int
func sideCombs(sides []int, leng int, startPos int, res []int) {
  if leng == 0 {
    // log.Println(res)
    allSideCombs = append(allSideCombs, copyArr(res))
    return
  }

  for i := startPos; i <= len(sides)-leng; i++ {
    // res = append(res, sides[i])
    res[len(res) - leng] = sides[i]
    sideCombs(sides, leng-1, i+1, res)
  }
}

func copyArr(src []int) []int {
  dest  := []int{}

  for _, i := range src {
    dest = append(dest, i)
  }

  return dest
}

