package main

import (
    "fmt"
)

func main() {
    numbers := []int{85, 10, 9, 8, 74, 37, 93, 89, 19, 56, 26, 63, 9, 35, 87, 24, 16, 27, 17, 55, 31, 7, 5, 45, 14, 38, 88, 2, 69, 4, 43, 46, 53, 32, 11, 86, 80, 47, 3, 71, 97, 49, 13, 56, 25, 66, 60, 98, 92, 18, 58, 56, 44, 44, 44, 44}
    err := detectRepeats(numbers)

    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Println("No >3 repeat numbers")
    }
}

// return an error if any number is found to have been repeated more than once in the input collection 
func detectRepeats(numbers []int) error {
    searched_nums := map[int]int{}

    for i, num := range numbers {
        _, searched := searched_nums[num]

        num_count := 0
        if !searched {
            for j := i; j < len(numbers); j++ {
                if numbers[j] == num {
                    num_count++
                }

                if num_count > 3 {
                    return fmt.Errorf("%d found >3 times.", numbers[j])
                }
            }
            searched_nums[num] = num
        }
    }
    return nil
}
