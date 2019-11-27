package main

import "fmt"

// Sum takes a fixed array and returns the sum of its content.
func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SumAll takes a variable number of slices, and sums the values.
func SumAll(arrays ...[]int) (sum []int) {
	for _, numbers := range arrays {
		sum = append(sum, Sum(numbers))
	}

	return
}

// SumAllTails takes the 'tail' of each array and sums the values.
func SumAllTails(arrays ...[]int) (sum []int) {
	for _, numbers := range arrays {
		if len(numbers) == 0 {
			sum = append(sum, 0)

		} else {
			tail := numbers[1:]
			sum = append(sum, Sum(tail))
		}
	}

	return
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(Sum(numbers))
}
